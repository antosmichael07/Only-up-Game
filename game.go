package main

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func init_game() (Player, Player, []rl.Rectangle, []SideLauncher, []Launcher, rl.Camera2D) {
	player := NewPlayer()
	player_2 := NewPlayer()

	collision_rects := []rl.Rectangle{
		rl.NewRectangle(25, 125, 250, 25),
		rl.NewRectangle(0, 400, 400, 25),
	}

	side_launchers := []SideLauncher{
		{rl.NewRectangle(250, 100, 25, 25), -8},
	}

	launchers := []Launcher{
		{rl.NewRectangle(25, 100, 25, 25), -8},
	}

	camera := rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)), rl.NewVector2(225, player.Position.Y-200), 0, 4)

	return player, player_2, collision_rects, side_launchers, launchers, camera
}

func game_loop() {
	player, player_2, collision_rects, jumpers, side_launchers, camera := init_game()

	var wg sync.WaitGroup
	wg.Add(1)
	go connection(&player, &player_2, &wg)
	wg.Wait()

	for !rl.WindowShouldClose() {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(camera)
		update_camera(&player, &camera)

		player.Input()
		player.Update(&collision_rects, &jumpers, &side_launchers)
		player_2.Update(&collision_rects, &jumpers, &side_launchers)

		for i := 0; i < len(collision_rects); i++ {
			rl.DrawRectangleRec(collision_rects[i], rl.Black)
		}
		for i := 0; i < len(jumpers); i++ {
			rl.DrawRectangleRec(jumpers[i].Rect, rl.Red)
		}
		for i := 0; i < len(side_launchers); i++ {
			rl.DrawRectangleRec(side_launchers[i].Rect, rl.Green)
		}

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func update_camera(player *Player, camera *rl.Camera2D) {
	camera.Target.Y = rl.Lerp(camera.Target.Y, player.Position.Y, 0.05*player.FrameTime)
}
