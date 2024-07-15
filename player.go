package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position     rl.Vector2
	Scale        rl.Vector2
	Texture      *rl.Texture2D
	Speed        float32
	Acceleration float32
	Gravity      float32
	GravityPower float32
	JumpPower    float32
}

func NewPlayer(texture *rl.Texture2D) Player {
	player := Player{}
	player.Position = rl.NewVector2(500, 100)
	player.Scale = rl.NewVector2(100, 100)
	player.Texture = texture
	player.Speed = 5.
	player.Gravity = 0.
	player.GravityPower = .15
	player.JumpPower = -5.

	return player
}

func (player *Player) Update(collision_rects *[]rl.Rectangle) {
	player.Move(collision_rects)
	player.Fall(collision_rects)
	player.Draw()
}

func (player *Player) Draw() {
	rl.DrawTexture(*player.Texture, int32(player.Position.X), int32(player.Position.Y), rl.White)
}

func (player *Player) Move(collision_rects *[]rl.Rectangle) {
	if rl.IsKeyDown(rl.KeyD) {
		player.Position.X += player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) {
		player.Position.X -= player.Speed
	}
	if rl.IsKeyDown(rl.KeyW) && player.on_ground(collision_rects) {
		player.Gravity = player.JumpPower
	}
}

func (player *Player) on_ground(collision_rects *[]rl.Rectangle) bool {
	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y+1, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			return true
		}
	}

	return false
}

func (player *Player) Fall(collision_rects *[]rl.Rectangle) {
	player.Gravity += player.GravityPower

	player_rect := rl.NewRectangle(player.Position.X, player.Position.Y+player.Gravity, player.Scale.X, player.Scale.Y)

	for i := 0; i < len(*collision_rects); i++ {
		if rl.CheckCollisionRecs(player_rect, (*collision_rects)[i]) {
			player.Gravity = 0
			player.Position.Y = (*collision_rects)[i].Y - player.Scale.Y
		}
	}

	player.Position.Y += player.Gravity
}
