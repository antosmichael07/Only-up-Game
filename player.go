package main

import (
	"time"

	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position           rl.Vector2
	Scale              rl.Vector2
	Speed              float32
	Acceleration       float32
	Gravity            float32
	GravityPower       float32
	JumpPower          float32
	Direction          int8
	SideLauncherPower  float32
	SideLauncherResist float32
	FrameTime          float32
	AnimationTimer     float32
	Keys               [3]byte
	Kicking            bool
}

type Launcher struct {
	Rect  rl.Rectangle
	Power float32
}

func NewPlayer() Player {
	player := Player{}
	player.Position = rl.NewVector2(100, 100)
	player.Scale = rl.NewVector2(21, 25)
	player.Speed = 2.
	player.Gravity = 0.
	player.GravityPower = .125
	player.JumpPower = -1.8
	player.Direction = 1
	player.SideLauncherPower = 0
	player.SideLauncherResist = .05
	player.FrameTime = 0
	player.AnimationTimer = 0

	return player
}

func (player *Player) Update(collision_rects *[]rl.Rectangle, side_launchers *[]SideLauncher, launchers *[]Launcher, player_textures *[][3]rl.Texture2D, players *[]Player, client *tcp.Client) {
	player.FrameTime = rl.GetFrameTime() * 60

	player.Movement(collision_rects)
	player.Fall(collision_rects)
	player.SideLauncher(side_launchers, collision_rects, client)
	player.Launcher(launchers)
	player.Drawing(player_textures)
}

func (player *Player) Drawing(player_textures *[][3]rl.Texture2D) {
	player.AnimationTimer += player.FrameTime

	if player.Direction < 0 {
		if player.Kicking {
			rl.DrawTexture((*player_textures)[4][0], int32(player.Position.X), int32(player.Position.Y), rl.White)
			go func() {
				time.Sleep(500 * time.Millisecond)
				player.Kicking = false
			}()
			return
		}
		if player.Gravity < 2.5 {
			player.Draw(&(*player_textures)[0])
		} else {
			player.Draw(&(*player_textures)[2])
		}
	} else {
		if player.Kicking {
			rl.DrawTexture((*player_textures)[4][1], int32(player.Position.X), int32(player.Position.Y), rl.White)
			go func() {
				time.Sleep(500 * time.Millisecond)
				player.Kicking = false
			}()
			return
		}
		if player.Gravity < 2.5 {
			player.Draw(&(*player_textures)[1])
		} else {
			player.Draw(&(*player_textures)[3])
		}
	}
}

func (player *Player) Draw(textures *[3]rl.Texture2D) {
	if player.Keys[0] != 1 && player.Keys[1] != 1 && player.Gravity == 0 && player.SideLauncherPower == 0 {
		rl.DrawTexture((*textures)[0], int32(player.Position.X), int32(player.Position.Y), rl.White)
		player.AnimationTimer = 0
		return
	}

	if player.AnimationTimer < 4 {
		rl.DrawTexture((*textures)[1], int32(player.Position.X), int32(player.Position.Y), rl.White)
	} else if player.AnimationTimer < 8 {
		rl.DrawTexture((*textures)[2], int32(player.Position.X), int32(player.Position.Y), rl.White)
	} else if player.AnimationTimer < 12 {
		rl.DrawTexture((*textures)[0], int32(player.Position.X), int32(player.Position.Y), rl.White)
	} else {
		rl.DrawTexture((*textures)[1], int32(player.Position.X), int32(player.Position.Y), rl.White)
		player.AnimationTimer -= 12
	}
}

func (player *Player) Movement(collision_rects *[]rl.Rectangle) {
	if player.Keys[0] == 1 {
		player.Direction = -1
		if player.SideLauncherPower == 0 {
			player.Move(collision_rects, player.Speed*float32(player.Direction))
		}
	} else if player.Keys[1] == 1 {
		player.Direction = 1
		if player.SideLauncherPower == 0 {
			player.Move(collision_rects, player.Speed*float32(player.Direction))
		}
	}

	if player.Keys[2] == 1 && player.OnGround(collision_rects) {
		player.Gravity = player.JumpPower
	}
}

func (player *Player) Move(collision_rects *[]rl.Rectangle, speed float32) {
	player_rect := rl.NewRectangle(player.Position.X+speed*player.FrameTime, player.Position.Y, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			if speed > 0 {
				player.Position.X = (*collision_rects)[i].X - player.Scale.X
			} else {
				player.Position.X = (*collision_rects)[i].X + (*collision_rects)[i].Width
			}
			return
		}
	}

	player.Position.X += speed * player.FrameTime
}

func (player *Player) Fall(collision_rects *[]rl.Rectangle) {
	player.Gravity += player.GravityPower * player.FrameTime

	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y+player.Gravity*player.FrameTime, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			player.Gravity = 0
			player.Position.Y = (*collision_rects)[i].Y - player.Scale.Y
			return
		}
	}

	player.Position.Y += player.Gravity * player.FrameTime
}

func (player *Player) OnGround(collision_rects *[]rl.Rectangle) bool {
	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y+1, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			return true
		}
	}

	return false
}

func (player *Player) SideLauncher(side_launchers *[]SideLauncher, collision_rects *[]rl.Rectangle, client *tcp.Client) {
	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*side_launchers); i++ {
		if rl.CheckCollisionRecs(player_rect, (*side_launchers)[i].Rect) && (*side_launchers)[i].AnimationTimer <= 0 {
			player.SideLauncherPower = (*side_launchers)[i].Power
			(*side_launchers)[i].AnimationTimer = 2
			client.SendData(event_side_launcher_launched, &[]byte{byte(i)})
			break
		}
	}

	if player.SideLauncherPower > 0 {
		if player.Keys[0] == 1 {
			player.SideLauncherPower -= player.SideLauncherResist * 5 * player.FrameTime
		} else if player.Keys[1] == 1 {
			player.SideLauncherPower -= player.SideLauncherResist / 5 * player.FrameTime
			if player.SideLauncherPower < player.Speed {
				player.SideLauncherPower = 0
			}
		} else {
			player.SideLauncherPower -= player.SideLauncherResist * player.FrameTime
		}

		if player.SideLauncherPower < 0 {
			player.SideLauncherPower = 0
		}
	} else if player.SideLauncherPower < 0 {
		if player.Keys[0] == 1 {
			player.SideLauncherPower += player.SideLauncherResist / 5 * player.FrameTime
			if player.SideLauncherPower > -player.Speed {
				player.SideLauncherPower = 0
			}
		} else if player.Keys[1] == 1 {
			player.SideLauncherPower += player.SideLauncherResist * 5 * player.FrameTime
		} else {
			player.SideLauncherPower += player.SideLauncherResist * player.FrameTime
		}

		if player.SideLauncherPower > 0 {
			player.SideLauncherPower = 0
		}
	}

	player_rect = rl.NewRectangle(player.Position.X+player.SideLauncherPower*player.FrameTime, player.Position.Y, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			if player.SideLauncherPower > 0 {
				player.SideLauncherPower = player.SideLauncherPower * -1
				player.Position.X = (*collision_rects)[i].X - player.Scale.X
			} else {
				player.SideLauncherPower = player.SideLauncherPower * -1
				player.Position.X = (*collision_rects)[i].X + (*collision_rects)[i].Width
			}
			return
		}
	}

	player.Position.X += player.SideLauncherPower * player.FrameTime
}

func (player *Player) Launcher(launchers *[]Launcher) {
	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*launchers); i++ {
		if rl.CheckCollisionRecs(player_rect, (*launchers)[i].Rect) {
			player.Gravity = (*launchers)[i].Power
			break
		}
	}
}

func (player *Player) Kick(players *[]Player, player_num *byte, client *tcp.Client, settings *Settings) {
	if rl.IsKeyPressed(settings.PlayerKick) {
		player_rect := rl.NewRectangle(player.Position.X, player.Position.Y, player.Scale.X, player.Scale.Y)

		for i := 0; i < len(*players); i++ {
			if i == int(*player_num) {
				continue
			}

			other_player_rect := rl.NewRectangle((*players)[i].Position.X, (*players)[i].Position.Y, (*players)[i].Scale.X, (*players)[i].Scale.Y)

			if rl.CheckCollisionRecs(player_rect, other_player_rect) {
				to_send := append([]byte{byte(i)}, float32_to_bytes((*players)[i].SideLauncherPower+(6*float32((*player).Direction)))...)
				to_send = append(to_send, byte(*player_num))
				client.SendData(event_player_kick, &to_send)
				break
			}
		}
	}
}

func (player *Player) DrawArrow(arrow *rl.Texture2D) {
	rl.DrawTexture(*arrow, int32(player.Position.X)+7, int32(player.Position.Y)-6, rl.White)
}
