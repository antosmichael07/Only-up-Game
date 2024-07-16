package main

import rl "github.com/gen2brain/raylib-go/raylib"

func init_game() (Player, []rl.Rectangle, []Jumper, rl.Camera2D) {
	player := NewPlayer()

	collision_rects := []rl.Rectangle{
		rl.NewRectangle(25, 125, 250, 25),
		rl.NewRectangle(0, 400, 400, 25),
	}

	jumpers := []Jumper{
		{rl.NewRectangle(25, 100, 25, 25), 1, 8},
		{rl.NewRectangle(250, 100, 25, 25), -1, 8},
	}

	camera := rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)), rl.NewVector2(225, player.Position.Y-200), 0, 4)

	return player, collision_rects, jumpers, camera
}

func game_loop() {
	player, collision_rects, jumpers, camera := init_game()

	for !rl.WindowShouldClose() {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(camera)
		update_camera(&player, &camera)

		player.Update(&collision_rects, &jumpers)

		for i := 0; i < len(collision_rects); i++ {
			rl.DrawRectangleRec(collision_rects[i], rl.Black)
		}
		for i := 0; i < len(jumpers); i++ {
			rl.DrawRectangleRec(jumpers[i].Rect, rl.Red)
		}

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func update_camera(player *Player, camera *rl.Camera2D) {
	camera.Target.Y = rl.Lerp(camera.Target.Y, player.Position.Y, 0.05*player.FrameTime)
}
