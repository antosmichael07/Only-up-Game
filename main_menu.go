package main

import (
	lgr "github.com/antosmichael07/Go-Logger"
	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main_menu(should_close_connection *bool) {
	for !rl.WindowShouldClose() {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.DrawText("Main Menu", 10, 10, 20, rl.Black)
		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeySpace) {
			client := tcp.NewClient("192.168.1.127:8080")
			client.Logger.Level = lgr.None
			err := client.Connect()
			if err != nil {
				continue
			}

			game_loop(should_close_connection, &client)
		}
	}
}
