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

	should_close_connection := false
	stop_trying_to_connect := false
	back_from_credits := false
	ip := ""
	init_buttons(&buttons, &input_box, &should_close_connection, &stop_trying_to_connect, &ip, &back_from_credits)

	main_menu(&buttons)
}
