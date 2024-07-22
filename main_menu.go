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

		buttons.Draw(0)

		rl.EndDrawing()
	}
}

func init_buttons(buttons *Buttons, input_box *rl.Texture2D, should_close_connection *bool, stop_trying_to_connect *bool, ip *string, back_from_credits *bool, player_textures *[][3]rl.Texture2D, arrow *rl.Texture2D, go_back *bool) {
	buttons.b_types[0].NewButton("join", int32(rl.GetScreenWidth()/2)-300, 100, "JOIN", 60, func(button *Button) {
		*stop_trying_to_connect = false

		for !*stop_trying_to_connect {
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

			if rl.IsKeyPressed(rl.KeyEscape) {
				*stop_trying_to_connect = true
			}

			if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
				*should_close_connection = false
				connect(ip, should_close_connection, player_textures, arrow, go_back, buttons)
			}
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

			if rl.IsKeyPressed(rl.KeyEscape) {
				*back_from_credits = true
			}
		}

		rl.BeginDrawing()
	})
	buttons.b_types[0].NewButton("quit", int32(rl.GetScreenWidth()/2)-300, int32(rl.GetScreenHeight())-250, "QUIT", 60, func(button *Button) {
		os.Exit(0)
	})

	buttons.b_types[1].NewButton("connect", int32(rl.GetScreenWidth()/2)-300, 400, "CONNECT", 60, func(button *Button) {
		rl.EndDrawing()
		*should_close_connection = false
		connect(ip, should_close_connection, player_textures, arrow, go_back, buttons)
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

	buttons.b_types[4].NewButton("back-from-error", int32(rl.GetScreenWidth()/2)-300, 750, "BACK", 60, func(button *Button) {
		*go_back = true
	})
}

func connect(ip *string, should_close_connection *bool, player_textures *[][3]rl.Texture2D, arrow *rl.Texture2D, go_back *bool, buttons *Buttons) {
	client := tcp.NewClient(*ip)
	client.Logger.Level = lgr.None

	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)
	rl.DrawText("LOADING...", int32(rl.GetScreenWidth()/2)-rl.MeasureText("LOADING...", 60)/2, 300, 60, rl.Black)
	rl.EndDrawing()

	err := client.Connect()
	if err != nil {
		*go_back = false

		for !*go_back {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			iterations := 0
			for i := 0; i < len(err.Error()); {
				last_space := i
				j := i
				for j < len(err.Error()) && err.Error()[j] != '\n' && rl.MeasureText(err.Error()[i:j+1], 60) < int32(rl.GetScreenWidth())-200 {
					if err.Error()[j] == ' ' {
						last_space = j
					}
					j++
				}

				rl.DrawText(err.Error()[i:last_space], int32(rl.GetScreenWidth())/2-rl.MeasureText(err.Error()[i:last_space], 60)/2, 100+int32(70*iterations), 60, rl.Black)

				i = last_space + 1
				iterations++
			}

			buttons.Draw(4)

			rl.EndDrawing()

			if rl.IsKeyPressed(rl.KeyEscape) {
				*go_back = true
			}
		}

		return
	}

	game_loop(should_close_connection, &client, player_textures, arrow)
}
