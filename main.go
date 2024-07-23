package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	init_window()

	buttons := NewButtons()

	button_normal := rl.LoadTexture("resources/textures/button_normal.png")
	button_focused := rl.LoadTexture("resources/textures/button_focused.png")
	button_pressed := rl.LoadTexture("resources/textures/button_pressed.png")
	clear_normal := rl.LoadTexture("resources/textures/clear_normal.png")
	clear_focused := rl.LoadTexture("resources/textures/clear_focused.png")
	clear_pressed := rl.LoadTexture("resources/textures/clear_pressed.png")
	input_box := rl.LoadTexture("resources/textures/input_box.png")

	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&clear_normal, &clear_focused, &clear_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)

	should_close_connection := false
	stop_trying_to_connect := false
	back_from_credits := false
	go_back := false
	ip := ""
	cursor := 0
	cursor_timer := float32(0)

	player_textures := [][3]rl.Texture2D{
		{
			rl.LoadTexture("resources/textures/player_00.png"),
			rl.LoadTexture("resources/textures/player_01.png"),
			rl.LoadTexture("resources/textures/player_02.png"),
		},
		{
			rl.LoadTexture("resources/textures/player_10.png"),
			rl.LoadTexture("resources/textures/player_11.png"),
			rl.LoadTexture("resources/textures/player_12.png"),
		},
		{
			rl.LoadTexture("resources/textures/player_scream_00.png"),
			rl.LoadTexture("resources/textures/player_scream_01.png"),
			rl.LoadTexture("resources/textures/player_scream_02.png"),
		},
		{
			rl.LoadTexture("resources/textures/player_scream_10.png"),
			rl.LoadTexture("resources/textures/player_scream_11.png"),
			rl.LoadTexture("resources/textures/player_scream_12.png"),
		},
		{
			rl.LoadTexture("resources/textures/player_kick_0.png"),
			rl.LoadTexture("resources/textures/player_kick_1.png"),
			rl.LoadTexture("resources/textures/player_kick_0.png"),
		},
	}

	arrow := rl.LoadTexture("resources/textures/arrow.png")

	init_buttons(&buttons, &input_box, &should_close_connection, &stop_trying_to_connect, &ip, &back_from_credits, &player_textures, &arrow, &go_back, &cursor, &cursor_timer)

	main_menu(&buttons)
}
