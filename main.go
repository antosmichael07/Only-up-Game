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
	little_button_normal := rl.LoadTexture("resources/textures/little_button_normal.png")
	little_button_focused := rl.LoadTexture("resources/textures/little_button_focused.png")
	little_button_pressed := rl.LoadTexture("resources/textures/little_button_pressed.png")
	clear_normal := rl.LoadTexture("resources/textures/clear_normal.png")
	clear_focused := rl.LoadTexture("resources/textures/clear_focused.png")
	clear_pressed := rl.LoadTexture("resources/textures/clear_pressed.png")
	input_box := rl.LoadTexture("resources/textures/input_box.png")

	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&clear_normal, &clear_focused, &clear_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&little_button_normal, &little_button_focused, &little_button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&little_button_normal, &little_button_focused, &little_button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)
	buttons.NewButtonType(&little_button_normal, &little_button_focused, &little_button_pressed)
	buttons.NewButtonType(&button_normal, &button_focused, &button_pressed)

	should_close_connection := false
	stop_trying_to_connect := false
	back_from_credits := false
	go_back := false
	ip := ""
	cursor := 0
	cursor_timer := float32(0)
	is_game_menu_open := false
	is_settings_open := false
	wait_for_server := false
	server_err := error(nil)
	err := error(nil)
	settings := load_settings()

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

	side_launcher_textures := [2][4]rl.Texture2D{
		{
			rl.LoadTexture("resources/textures/side_launcher_00.png"),
			rl.LoadTexture("resources/textures/side_launcher_01.png"),
			rl.LoadTexture("resources/textures/side_launcher_02.png"),
			rl.LoadTexture("resources/textures/side_launcher_03.png"),
		},
		{
			rl.LoadTexture("resources/textures/side_launcher_10.png"),
			rl.LoadTexture("resources/textures/side_launcher_11.png"),
			rl.LoadTexture("resources/textures/side_launcher_12.png"),
			rl.LoadTexture("resources/textures/side_launcher_13.png"),
		},
	}

	launcher_texture := rl.LoadTexture("resources/textures/launcher.png")

	arrow := rl.LoadTexture("resources/textures/arrow.png")

	rl.InitAudioDevice()
	button_click_sound := rl.LoadSound("resources/sounds/button_click.wav")

	background_texture := rl.LoadTexture("resources/textures/background.png")

	pre_objects[OBJECT_CONTAINER_RED].Texture = rl.LoadTexture("./resources/textures/objects/container_red.png")
	pre_objects[OBJECT_CONTAINER_GREEN].Texture = rl.LoadTexture("./resources/textures/objects/container_green.png")
	pre_objects[OBJECT_CONTAINER_BLUE].Texture = rl.LoadTexture("./resources/textures/objects/container_blue.png")
	pre_objects[OBJECT_CONTAINER_YELLOW].Texture = rl.LoadTexture("./resources/textures/objects/container_yellow.png")
	pre_objects[OBJECT_CARDBOARD_BOX].Texture = rl.LoadTexture("./resources/textures/objects/cardboard_box.png")
	pre_objects[OBJECT_PALLETS].Texture = rl.LoadTexture("./resources/textures/objects/pallets.png")
	pre_objects[OBJECT_METAL_PIPE].Texture = rl.LoadTexture("./resources/textures/objects/metal_pipe.png")
	pre_objects[OBJECT_METAL_SUPPORT].Texture = rl.LoadTexture("./resources/textures/objects/metal_support.png")
	pre_objects[OBJECT_METAL_SUPPORT_HORIZONTAL].Texture = rl.LoadTexture("./resources/textures/objects/metal_support_horizontal.png")
	pre_objects[OBJECT_BRICK].Texture = rl.LoadTexture("./resources/textures/objects/brick.png")
	pre_objects[OBJECT_BOLT].Texture = rl.LoadTexture("./resources/textures/objects/bolt.png")
	pre_objects[OBJECT_TOILET].Texture = rl.LoadTexture("./resources/textures/objects/toilet.png")

	buttons.SetSound(&button_click_sound)
	init_buttons(&buttons, &input_box, &should_close_connection, &stop_trying_to_connect, &ip, &back_from_credits, &player_textures, &arrow, &go_back, &cursor, &cursor_timer, &is_game_menu_open, &err, &side_launcher_textures, &settings, &is_settings_open, &wait_for_server, &server_err, &launcher_texture, &background_texture)

	main_menu(&buttons)
}
