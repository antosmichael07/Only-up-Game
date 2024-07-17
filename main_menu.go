package main

import (
	"os"
	"os/exec"

	lgr "github.com/antosmichael07/Go-Logger"
	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main_menu(buttons *Buttons) {
	for {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		(*buttons).Draw(0)

		rl.EndDrawing()
	}
}

func init_buttons(buttons *Buttons, input_box *rl.Texture2D, should_close_connection *bool, stop_trying_to_connect *bool, ip *string, back_from_credits *bool) {
	buttons.b_types[0].NewButton("join", int32(rl.GetScreenWidth()/2)-300, 100, "JOIN", 60, func(button *Button) {
		*stop_trying_to_connect = false

		for !rl.IsKeyPressed(rl.KeyEnter) && !rl.IsKeyPressed(rl.KeyKpEnter) && !rl.IsKeyPressed(rl.KeyEscape) && !*stop_trying_to_connect {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			char := rl.GetCharPressed()
			if char != 0 {
				*ip += string(char)
			}

			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyV) {
				*ip += rl.GetClipboardText()
			}

			rl.DrawTexture(*input_box, int32(rl.GetScreenWidth()/2)-500, 100, rl.White)
			rl.DrawText(*ip, int32(rl.GetScreenWidth()/2)-rl.MeasureText(*ip, 60)/2, 145, 60, rl.Black)

			buttons.Draw(1)
			buttons.Draw(2)

			rl.EndDrawing()

			if rl.IsKeyPressed(rl.KeyBackspace) {
				if len(*ip) > 0 {
					*ip = (*ip)[:len(*ip)-1]
				}
			}
		}

		if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			connect(ip, should_close_connection)
			*stop_trying_to_connect = true
		}
	})
	buttons.b_types[0].NewButton("credits", int32(rl.GetScreenWidth()/2)-300, 300, "CREDITS", 60, func(button *Button) {
		*back_from_credits = false
		rl.EndDrawing()

		for !*back_from_credits {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			rl.DrawTexture(*input_box, int32(rl.GetScreenWidth()/2)-500, 100, rl.White)
			rl.DrawText("Made By Mispul", int32(rl.GetScreenWidth()/2)-rl.MeasureText("Made By Mispul", 60)/2, 145, 60, rl.Black)

			buttons.Draw(3)

			rl.EndDrawing()
		}

		rl.BeginDrawing()
	})
	buttons.b_types[0].NewButton("quit", int32(rl.GetScreenWidth()/2)-300, int32(rl.GetScreenHeight())-250, "QUIT", 60, func(button *Button) {
		os.Exit(0)
	})

	buttons.b_types[1].NewButton("connect", int32(rl.GetScreenWidth()/2)-300, 400, "CONNECT", 60, func(button *Button) {
		rl.EndDrawing()
		connect(ip, should_close_connection)
		*stop_trying_to_connect = true
		rl.BeginDrawing()
	})
	buttons.b_types[1].NewButton("back-from-connecting", int32(rl.GetScreenWidth()/2)-300, 600, "BACK", 60, func(button *Button) {
		*stop_trying_to_connect = true
	})

	buttons.b_types[2].NewButton("clear-input-box", int32(rl.GetScreenWidth()/2)+550, 100, "", 60, func(button *Button) {
		*ip = ""
	})

	buttons.b_types[3].NewButton("github", int32(rl.GetScreenWidth()/2)-300, 400, "GITHUB", 60, func(button *Button) {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://github.com/antosmichael07").Start()
	})

	buttons.b_types[3].NewButton("back-from-credits", int32(rl.GetScreenWidth()/2)-300, 600, "BACK", 60, func(button *Button) {
		*back_from_credits = true
	})
}

func connect(ip *string, should_close_connection *bool) {
	client := tcp.NewClient(*ip)
	client.Logger.Level = lgr.None

	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)
	rl.DrawText("LOADING...", int32(rl.GetScreenWidth()/2)-rl.MeasureText("LOADING...", 60)/2, 300, 60, rl.Black)
	rl.EndDrawing()

	err := client.Connect()
	if err != nil {
		return
	}

	game_loop(should_close_connection, &client)
}
