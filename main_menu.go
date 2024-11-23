package main

import (
	"fmt"
	"os"
	"os/exec"

	lgr "github.com/antosmichael07/Go-Logger"
	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main_menu(buttons *Buttons) {

	if !rl.IsAudioDeviceReady() {
		for {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			msg := "audio device didn't initialize correctly, please restart the game"

			iterations := 0
			for i := 0; i < len(msg); {
				last_space := i

				for j := i; msg[j] != '\n' && rl.MeasureText(msg[i:j+1], 60) < int32(rl.GetScreenWidth())-200; j++ {
					if j+1 == len(msg) {
						last_space = j + 1
						break
					}
					if msg[j] == ' ' {
						last_space = j
					}
				}

				rl.DrawText(msg[i:last_space], int32(rl.GetScreenWidth())/2-rl.MeasureText(msg[i:last_space], 60)/2, 100+int32(70*iterations), 60, rl.Black)

				i = last_space + 1
				iterations++
			}

			buttons.Draw(11)

			rl.EndDrawing()
		}
	}

	for {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		buttons.Draw(0)
		buttons.Draw(8)

		rl.EndDrawing()
	}
}

func init_buttons(buttons *Buttons, input_box *rl.Texture2D, should_close_connection *bool, stop_trying_to_connect *bool, ip *string, back_from_credits *bool, player_textures *[][3]rl.Texture2D, arrow *rl.Texture2D, go_back *bool, cursor *int, cursor_timer *float32, is_game_menu_open *bool, err *error, side_launcher_textures *[2][4]rl.Texture2D, settings *Settings, is_settings_open *bool, wait_for_server *bool, server_err *error, launcher_texture *rl.Texture2D, background_texture *rl.Texture2D) {
	buttons.b_types[0].NewButton("join", int32(rl.GetScreenWidth()/2)-300, 100, "JOIN", 60, func(button *Button) {
		*cursor_timer = 0.
		*stop_trying_to_connect = false
		*cursor = len(*ip)

		for !*stop_trying_to_connect {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			if (rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight)) && *cursor < len(*ip) {
				*cursor_timer = 0
				*cursor++
			}
			if (rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft)) && *cursor > 0 {
				*cursor_timer = 0
				*cursor--
			}

			char := rl.GetCharPressed()
			if char != 0 {
				*cursor_timer = 0
				*ip = fmt.Sprintf("%s%s%s", (*ip)[:*cursor], string(char), (*ip)[*cursor:])
				*cursor++
			}

			if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyV) {
				*cursor_timer = 0
				*ip = fmt.Sprintf("%s%s%s", (*ip)[:*cursor], rl.GetClipboardText(), (*ip)[*cursor:])
			}

			rl.DrawTexture(*input_box, int32(rl.GetScreenWidth()/2)-500, 100, rl.White)
			rl.DrawText(*ip, int32(rl.GetScreenWidth()/2)-rl.MeasureText(*ip, 60)/2, 145, 60, rl.Black)
			if *cursor_timer < .5 {
				rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-rl.MeasureText(*ip, 60)/2+rl.MeasureText((*ip)[:*cursor], 60)+2, 145, 2, 60, rl.Black)
			}

			buttons.Draw(1)
			buttons.Draw(2)

			rl.EndDrawing()

			if rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace) {
				if *cursor > 0 {
					*cursor_timer = 0
					*ip = fmt.Sprintf("%s%s", (*ip)[:*cursor-1], (*ip)[*cursor:])
					*cursor--
				}
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				*stop_trying_to_connect = true
			}

			if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
				*should_close_connection = false
				connect(ip, should_close_connection, player_textures, arrow, go_back, buttons, is_game_menu_open, err, side_launcher_textures, settings, launcher_texture, background_texture)
			}

			*cursor_timer += rl.GetFrameTime()
			*cursor_timer -= float32(int(*cursor_timer))
		}
	})

	buttons.b_types[0].NewButton("open-settings", int32(rl.GetScreenWidth()/2)-300, 300, "SETTINGS", 60, func(button *Button) {
		*is_settings_open = true

		rl.EndDrawing()

		for *is_settings_open {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			rl.DrawText("LEFT", int32(rl.GetScreenWidth()/2)-rl.MeasureText("LEFT", 60)/2-325, 130, 60, rl.Black)
			rl.DrawText("RIGHT", int32(rl.GetScreenWidth()/2)-rl.MeasureText("RIGHT", 60)/2+325, 130, 60, rl.Black)
			rl.DrawText("JUMP", int32(rl.GetScreenWidth()/2)-rl.MeasureText("JUMP", 60)/2, 55, 60, rl.Black)
			rl.DrawText("KICK", int32(rl.GetScreenWidth()/2)-rl.MeasureText("KICK", 60)/2, 330, 60, rl.Black)
			buttons.Draw(6)
			buttons.Draw(7)
			buttons.Draw(10)

			rl.EndDrawing()

			if rl.IsKeyPressed(rl.KeyEscape) {
				*is_settings_open = false
			}
		}

		rl.BeginDrawing()
	})

	buttons.b_types[0].NewButton("exit", int32(rl.GetScreenWidth()/2)-300, int32(rl.GetScreenHeight())-250, "EXIT", 60, func(button *Button) {
		os.Exit(0)
	})

	buttons.b_types[1].NewButton("connect", int32(rl.GetScreenWidth()/2)-300, 400, "CONNECT", 60, func(button *Button) {
		rl.EndDrawing()
		*should_close_connection = false
		connect(ip, should_close_connection, player_textures, arrow, go_back, buttons, is_game_menu_open, err, side_launcher_textures, settings, launcher_texture, background_texture)
		rl.BeginDrawing()
	})

	buttons.b_types[1].NewButton("back-from-connecting", int32(rl.GetScreenWidth()/2)-300, 600, "BACK", 60, func(button *Button) {
		*stop_trying_to_connect = true
	})

	buttons.b_types[2].NewButton("clear-input-box", int32(rl.GetScreenWidth()/2)+550, 100, "", 60, func(button *Button) {
		*ip = ""
		*cursor_timer = 0
		*cursor = 0
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

	buttons.b_types[4].NewButton("copy-error", int32(rl.GetScreenWidth()/2)-300, 550, "COPY ERROR", 60, func(button *Button) {
		rl.SetClipboardText((*err).Error())
	})

	buttons.b_types[5].NewButton("back-from-game-menu", int32(rl.GetScreenWidth()/2)-300, 100, "BACK TO GAME", 60, func(button *Button) {
		*is_game_menu_open = false
	})

	buttons.b_types[5].NewButton("quit-from-server", int32(rl.GetScreenWidth()/2)-300, int32(rl.GetScreenHeight())-250, "QUIT", 60, func(button *Button) {
		*should_close_connection = true
		*is_game_menu_open = false
	})

	buttons.b_types[6].NewButton("set-player-left-setting", int32(rl.GetScreenWidth()/2)-475, 200, string(settings.PlayerLeft), 60, func(button *Button) {
		rl.EndDrawing()
		set_control_setting(&settings.PlayerLeft, "left", input_box, buttons, settings)
		rl.BeginDrawing()
	})

	buttons.b_types[6].NewButton("set-player-right-setting", int32(rl.GetScreenWidth()/2)+175, 200, string(settings.PlayerRight), 60, func(button *Button) {
		rl.EndDrawing()
		set_control_setting(&settings.PlayerRight, "right", input_box, buttons, settings)
		rl.BeginDrawing()
	})

	buttons.b_types[6].NewButton("set-player-jump-setting", int32(rl.GetScreenWidth()/2)-150, 125, string(settings.PlayerJump), 60, func(button *Button) {
		rl.EndDrawing()
		set_control_setting(&settings.PlayerJump, "jump", input_box, buttons, settings)
		rl.BeginDrawing()
	})

	buttons.b_types[6].NewButton("set-player-kick-setting", int32(rl.GetScreenWidth()/2)-150, 400, string(settings.PlayerKick), 60, func(button *Button) {
		rl.EndDrawing()
		set_control_setting(&settings.PlayerKick, "kick", input_box, buttons, settings)
		rl.BeginDrawing()
	})

	buttons.b_types[7].NewButton("back-from-settings", int32(rl.GetScreenWidth()/2)-300, int32(rl.GetScreenHeight())-250, "BACK", 60, func(button *Button) {
		*is_settings_open = false
	})

	buttons.b_types[8].NewButton("open-credits", int32(rl.GetScreenWidth())-400, int32(rl.GetScreenHeight())-250, "CREDITS", 50, func(button *Button) {
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

	buttons.b_types[9].NewButton("copy-server-error", int32(rl.GetScreenWidth())/2-300, int32(rl.GetScreenHeight())-450, "COPY ERROR", 60, func(button *Button) {
		rl.SetClipboardText((*server_err).Error())
	})

	buttons.b_types[9].NewButton("back-from-server-error", int32(rl.GetScreenWidth())/2-300, int32(rl.GetScreenHeight())-250, "BACK", 60, func(button *Button) {
		*wait_for_server = false
	})

	buttons.b_types[10].NewButton("set-default-settings", int32(rl.GetScreenWidth())-400, int32(rl.GetScreenHeight())-250, "DEFAULT", 50, func(button *Button) {
		*settings = default_settings()
		save_settings(settings)
		buttons.b_types[6].SetText("set-player-left-setting", "A")
		buttons.b_types[6].SetText("set-player-right-setting", "D")
		buttons.b_types[6].SetText("set-player-jump-setting", "W")
		buttons.b_types[6].SetText("set-player-kick-setting", "SPACE")
	})

	buttons.b_types[11].NewButton("close-err-audio", int32(rl.GetScreenWidth())/2-300, int32(rl.GetScreenHeight())-250, "EXIT", 60, func(button *Button) {
		os.Exit(0)
	})

	if settings.PlayerLeft == rl.KeySpace {
		buttons.b_types[6].SetText("set-player-left-setting", "SPACE")
	}
	if settings.PlayerRight == rl.KeySpace {
		buttons.b_types[6].SetText("set-player-right-setting", "SPACE")
	}
	if settings.PlayerJump == rl.KeySpace {
		buttons.b_types[6].SetText("set-player-jump-setting", "SPACE")
	}
	if settings.PlayerKick == rl.KeySpace {
		buttons.b_types[6].SetText("set-player-kick-setting", "SPACE")
	}
}

func connect(ip *string, should_close_connection *bool, player_textures *[][3]rl.Texture2D, arrow *rl.Texture2D, go_back *bool, buttons *Buttons, is_game_menu_open *bool, err *error, side_launcher_textures *[2][4]rl.Texture2D, settings *Settings, launcher_texture *rl.Texture2D, background_texture *rl.Texture2D) {
	client := tcp.NewClient(*ip)
	client.Logger.Level = lgr.None

	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)
	rl.DrawText("CONNECTING...", int32(rl.GetScreenWidth()/2)-rl.MeasureText("CONNECTING...", 60)/2, 300, 60, rl.Black)
	rl.EndDrawing()

	*err = client.Connect()
	if *err != nil {
		*go_back = false

		for !*go_back {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			iterations := 0
			for i := 0; i < len((*err).Error()); {
				last_space := i

				for j := i; (*err).Error()[j] != '\n' && rl.MeasureText((*err).Error()[i:j+1], 60) < int32(rl.GetScreenWidth())-200; j++ {
					if j+1 == len((*err).Error()) {
						last_space = j + 1
						break
					}
					if (*err).Error()[j] == ' ' {
						last_space = j
					}
				}

				rl.DrawText((*err).Error()[i:last_space], int32(rl.GetScreenWidth())/2-rl.MeasureText((*err).Error()[i:last_space], 60)/2, 100+int32(70*iterations), 60, rl.Black)

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

	game_loop(should_close_connection, &client, player_textures, arrow, buttons, is_game_menu_open, side_launcher_textures, err, settings, launcher_texture, background_texture)

	if *err != nil {
		*go_back = false

		for !*go_back {
			window_manager()
			rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)

			iterations := 0
			for i := 0; i < len((*err).Error()); {
				last_space := i

				for j := i; (*err).Error()[j] != '\n' && rl.MeasureText((*err).Error()[i:j+1], 60) < int32(rl.GetScreenWidth())-200; j++ {
					if j+1 == len((*err).Error()) {
						last_space = j + 1
						break
					}
					if (*err).Error()[j] == ' ' {
						last_space = j
					}
				}

				rl.DrawText((*err).Error()[i:last_space], int32(rl.GetScreenWidth())/2-rl.MeasureText((*err).Error()[i:last_space], 60)/2, 100+int32(70*iterations), 60, rl.Black)

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
}
