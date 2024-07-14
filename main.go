package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	init_window()
	game_loop()

	rl.CloseWindow()
}
