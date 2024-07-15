package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position       rl.Vector2
	Scale          rl.Vector2
	Texture0       rl.Texture2D
	Texture1       rl.Texture2D
	TextureScream0 rl.Texture2D
	TextureScream1 rl.Texture2D
	Speed          float32
	Acceleration   float32
	Gravity        float32
	GravityPower   float32
	JumpPower      float32
	Direction      uint8
}

func NewPlayer() Player {
	player := Player{}
	player.Position = rl.NewVector2(500, 100)
	player.Scale = rl.NewVector2(84, 100)
	player.Texture0 = rl.LoadTexture("resources/textures/player_0.png")
	player.Texture1 = rl.LoadTexture("resources/textures/player_1.png")
	player.TextureScream0 = rl.LoadTexture("resources/textures/player_scream_0.png")
	player.TextureScream1 = rl.LoadTexture("resources/textures/player_scream_1.png")
	player.Speed = 2.5
	player.Gravity = 0.
	player.GravityPower = .15
	player.JumpPower = -5.
	player.Direction = 0

	return player
}

func (player *Player) Update(collision_rects *[]rl.Rectangle) {
	player.Movement(collision_rects)
	player.Fall(collision_rects)
	player.Draw()
}

func (player *Player) Draw() {
	if player.Direction == 0 {
		if player.Gravity < 5 {
			rl.DrawTexture(player.Texture0, int32(player.Position.X), int32(player.Position.Y), rl.White)
		} else {
			rl.DrawTexture(player.TextureScream0, int32(player.Position.X), int32(player.Position.Y), rl.White)
		}
	} else {
		if player.Gravity < 5 {
			rl.DrawTexture(player.Texture1, int32(player.Position.X), int32(player.Position.Y), rl.White)
		} else {
			rl.DrawTexture(player.TextureScream1, int32(player.Position.X), int32(player.Position.Y), rl.White)
		}
	}
}

func (player *Player) Movement(collision_rects *[]rl.Rectangle) {
	if rl.IsKeyDown(rl.KeyA) {
		player.Move(collision_rects, -player.Speed)
		player.Direction = 0
	}
	if rl.IsKeyDown(rl.KeyD) {
		player.Move(collision_rects, player.Speed)
		player.Direction = 1
	}

	if rl.IsKeyDown(rl.KeyW) && player.OnGround(collision_rects) {
		player.Gravity = player.JumpPower
	}
}

func (player *Player) Move(collision_rects *[]rl.Rectangle, speed float32) {
	player_rect := rl.NewRectangle(player.Position.X+speed, player.Position.Y, player.Scale.X, player.Scale.Y)

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

	player.Position.X += speed
}

func (player *Player) Fall(collision_rects *[]rl.Rectangle) {
	player.Gravity += player.GravityPower

	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y+player.Gravity, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			player.Gravity = 0
			player.Position.Y = (*collision_rects)[i].Y - player.Scale.Y
			return
		}
	}

	player.Position.Y += player.Gravity
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
