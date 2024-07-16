package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position       rl.Vector2
	Scale          rl.Vector2
	Texture0       [3]rl.Texture2D
	Texture1       [3]rl.Texture2D
	TextureScream0 [3]rl.Texture2D
	TextureScream1 [3]rl.Texture2D
	Speed          float32
	Acceleration   float32
	Gravity        float32
	GravityPower   float32
	JumpPower      float32
	Direction      int8
	JumperPower    float32
	JumperResist   float32
	FrameTime      float32
	AnimationTimer float32
}

type Jumper struct {
	Rect      rl.Rectangle
	Direction int8
	Power     float32
}

func NewPlayer() Player {
	player := Player{}
	player.Position = rl.NewVector2(100, 100)
	player.Scale = rl.NewVector2(21, 25)
	player.Texture0[0] = rl.LoadTexture("resources/textures/player_00.png")
	player.Texture0[1] = rl.LoadTexture("resources/textures/player_01.png")
	player.Texture0[2] = rl.LoadTexture("resources/textures/player_02.png")
	player.Texture1[0] = rl.LoadTexture("resources/textures/player_10.png")
	player.Texture1[1] = rl.LoadTexture("resources/textures/player_11.png")
	player.Texture1[2] = rl.LoadTexture("resources/textures/player_12.png")
	player.TextureScream0[0] = rl.LoadTexture("resources/textures/player_scream_00.png")
	player.TextureScream0[1] = rl.LoadTexture("resources/textures/player_scream_01.png")
	player.TextureScream0[2] = rl.LoadTexture("resources/textures/player_scream_02.png")
	player.TextureScream1[0] = rl.LoadTexture("resources/textures/player_scream_10.png")
	player.TextureScream1[1] = rl.LoadTexture("resources/textures/player_scream_11.png")
	player.TextureScream1[2] = rl.LoadTexture("resources/textures/player_scream_12.png")
	player.Speed = 2.
	player.Gravity = 0.
	player.GravityPower = .125
	player.JumpPower = -1.8
	player.Direction = 1
	player.JumperPower = 0
	player.JumperResist = .05
	player.FrameTime = 0
	player.AnimationTimer = 0

	return player
}

func (player *Player) Update(collision_rects *[]rl.Rectangle, jumpers *[]Jumper) {
	player.FrameTime = rl.GetFrameTime() * 60

	player.Movement(collision_rects)
	player.Fall(collision_rects)
	player.Jumper(jumpers)
	player.Drawing()
}

func (player *Player) Drawing() {
	player.AnimationTimer += player.FrameTime

	if player.Direction < 0 {
		if player.Gravity < 2.5 {
			player.Draw(&player.Texture0)
		} else {
			player.Draw(&player.TextureScream0)
		}
	} else {
		if player.Gravity < 2.5 {
			player.Draw(&player.Texture1)
		} else {
			player.Draw(&player.TextureScream1)
		}
	}
}

func (player *Player) Draw(textures *[3]rl.Texture2D) {
	if !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyD) && player.Gravity == 0 && player.JumperPower == 0 {
		rl.DrawTexture((*textures)[0], int32(player.Position.X), int32(player.Position.Y), rl.White)
		player.AnimationTimer = 0
		return
	}

	if player.AnimationTimer < 5 {
		rl.DrawTexture((*textures)[0], int32(player.Position.X), int32(player.Position.Y), rl.White)
	} else if player.AnimationTimer < 10 {
		rl.DrawTexture((*textures)[1], int32(player.Position.X), int32(player.Position.Y), rl.White)
	} else if player.AnimationTimer < 15 {
		rl.DrawTexture((*textures)[2], int32(player.Position.X), int32(player.Position.Y), rl.White)
	} else {
		rl.DrawTexture((*textures)[0], int32(player.Position.X), int32(player.Position.Y), rl.White)
		player.AnimationTimer = 0
	}
}

func (player *Player) Movement(collision_rects *[]rl.Rectangle) {
	if rl.IsKeyDown(rl.KeyA) {
		player.Direction = -1
		if player.JumperPower == 0 {
			player.Move(collision_rects, player.Speed*float32(player.Direction))
		}
	} else if rl.IsKeyDown(rl.KeyD) {
		player.Direction = 1
		if player.JumperPower == 0 {
			player.Move(collision_rects, player.Speed*float32(player.Direction))
		}
	}

	if rl.IsKeyDown(rl.KeyW) && player.OnGround(collision_rects) {
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

func (player Player) OnGround(collision_rects *[]rl.Rectangle) bool {
	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y+1, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			return true
		}
	}

	return false
}

func (player *Player) Jumper(jumpers *[]Jumper) {
	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*jumpers); i++ {
		if rl.CheckCollisionRecs(player_rect, (*jumpers)[i].Rect) {
			player.JumperPower = (*jumpers)[i].Power * float32((*jumpers)[i].Direction)
			break
		}
	}

	if player.JumperPower > 0 {
		if rl.IsKeyDown(rl.KeyA) {
			player.JumperPower -= player.JumperResist * 2
		} else if rl.IsKeyDown(rl.KeyD) {
			player.JumperPower -= player.JumperResist / 2
			if player.JumperPower < player.Speed {
				player.JumperPower = 0
			}
		} else {
			player.JumperPower -= player.JumperResist
		}

		if player.JumperPower < 0 {
			player.JumperPower = 0
		}
	} else if player.JumperPower < 0 {
		if rl.IsKeyDown(rl.KeyA) {
			player.JumperPower += player.JumperResist / 2
			if player.JumperPower > -player.Speed {
				player.JumperPower = 0
			}
		} else if rl.IsKeyDown(rl.KeyD) {
			player.JumperPower += player.JumperResist * 2
		} else {
			player.JumperPower += player.JumperResist
		}

		if player.JumperPower > 0 {
			player.JumperPower = 0
		}
	}

	player.Position.X += player.JumperPower * player.FrameTime
}
