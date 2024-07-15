package main

import rl "github.com/gen2brain/raylib-go/raylib"

func init_game() (Player, []rl.Rectangle, rl.Camera2D) {
	player := NewPlayer()

	collision_rects := []rl.Rectangle{
		rl.NewRectangle(100, 500, 1000, 100),
		rl.NewRectangle(0, 1080, 1920, 100),
	}

	camera := rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)), rl.NewVector2(800, player.Position.Y-200), 0, 1)

	return player, collision_rects, camera
}

func game_loop() {
	player, collision_rects, camera := init_game()

	for !rl.WindowShouldClose() {
		fps_reducer()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(camera)
		update_camera(&player, &camera)

		player.Update(&collision_rects)

		for i := 0; i < len(collision_rects); i++ {
			rl.DrawRectangleRec(collision_rects[i], rl.Black)
		}

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func update_camera(player *Player, camera *rl.Camera2D) {
	camera.Target.Y = player.Position.Y - 200
}
