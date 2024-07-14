package main

import rl "github.com/gen2brain/raylib-go/raylib"

func game_loop() {
	for !rl.WindowShouldClose() {
		fps_reducer()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
