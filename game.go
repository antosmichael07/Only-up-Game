package main

import rl "github.com/gen2brain/raylib-go/raylib"

func init_game() (Player, []rl.Rectangle) {
	player_texture := rl.LoadTexture("resources/textures/player.png")
	player := NewPlayer(&player_texture)

	collision_rects := []rl.Rectangle{
		rl.NewRectangle(100, 500, 1000, 100),
		rl.NewRectangle(1300, 500, 500, 100),
		rl.NewRectangle(800, 400, 100, 100),
	}

	return player, collision_rects
}

func game_loop() {
	player, collision_rects := init_game()

	for !rl.WindowShouldClose() {
		fps_reducer()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		player.Update(&collision_rects)

		for i := 0; i < len(collision_rects); i++ {
			rl.DrawRectangleRec(collision_rects[i], rl.Black)
		}

		rl.EndDrawing()
	}
}
