package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Buttons struct {
	sound          *rl.Sound
	pressed_button bool
	b_types        []ButtonType
}

type ButtonType struct {
	normal  *rl.Texture2D
	focused *rl.Texture2D
	pressed *rl.Texture2D
	buttons []Button
}

type Button struct {
	id                          string
	position                    [2]int32
	text                        string
	font_size                   int32
	function                    func(*Button)
	precalculated_text_position [2]int32
}

func NewButtons() Buttons {
	return Buttons{&rl.Sound{}, false, []ButtonType{}}
}

func (b *Buttons) NewButtonType(normal *rl.Texture2D, focused *rl.Texture2D, pressed *rl.Texture2D) {
	button_type := ButtonType{normal, focused, pressed, []Button{}}
	b.b_types = append(b.b_types, button_type)
}

func (b *ButtonType) NewButton(id string, position_x int32, position_y int32, text string, font_size int32, function func(*Button)) {
	button := Button{id, [2]int32{position_x, position_y}, text, font_size, function, [2]int32{0, 0}}
	b.buttons = append(b.buttons, button)
	b.RecalculateTextPosition(len(b.buttons) - 1)
}

func (b *ButtonType) NewButtonSelect(id string, position_x int32, position_y int32, font_size int32, options []string) {
	function := func(button *Button) {
		is_selected := false
		for i := 0; i < len(b.buttons); i++ {
			if b.buttons[i].id == id+"_select" || b.buttons[i].id == id+"_select_now" {
				is_selected = true
			}
		}

		b.DeleteButtonType(id + "_select")
		b.DeleteButtonType(id + "_select_now")

		if !is_selected {
			for j := 0; j < len(options); j++ {
				b.NewButton(id+"_select_now", position_x, position_y+int32(j+1)*50, options[j], font_size, func(button *Button) {
					b.DeleteButtonType(id + "_select")
					b.SetText(id, options[j])
				})
			}
		}
	}
	button := Button{id, [2]int32{position_x, position_y}, "Select", font_size, function, [2]int32{0, 0}}
	b.buttons = append(b.buttons, button)
	b.RecalculateTextPosition(len(b.buttons) - 1)
}

func (b *Buttons) Draw(button_type int) {
	for i := 0; i < len(b.b_types[button_type].buttons); i++ {
		if len(b.b_types[button_type].buttons[i].id) > 10 && b.b_types[button_type].buttons[i].id[len(b.b_types[button_type].buttons[i].id)-11:] == "_select_now" {
			b.b_types[button_type].buttons[i].id = b.b_types[button_type].buttons[i].id[:len(b.b_types[button_type].buttons[i].id)-4]
		}
	}
	b.pressed_button = false
	for i := 0; i < len(b.b_types[button_type].buttons); i++ {
		if len(b.b_types[button_type].buttons[i].id) < 7 || b.b_types[button_type].buttons[i].id[len(b.b_types[button_type].buttons[i].id)-7:] != "_select" {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(b.b_types[button_type].buttons[i].position[0]), float32(b.b_types[button_type].buttons[i].position[1]), float32(b.b_types[button_type].normal.Width), float32(b.b_types[button_type].normal.Height))) {
				if rl.IsMouseButtonDown(rl.MouseLeftButton) {
					rl.DrawTexture(*b.b_types[button_type].pressed, b.b_types[button_type].buttons[i].position[0], b.b_types[button_type].buttons[i].position[1], rl.White)
				} else {
					rl.DrawTexture(*b.b_types[button_type].focused, b.b_types[button_type].buttons[i].position[0], b.b_types[button_type].buttons[i].position[1], rl.White)
				}
			} else {
				rl.DrawTexture(*b.b_types[button_type].normal, b.b_types[button_type].buttons[i].position[0], b.b_types[button_type].buttons[i].position[1], rl.White)
			}
			rl.DrawText(b.b_types[button_type].buttons[i].text, b.b_types[button_type].buttons[i].precalculated_text_position[0], b.b_types[button_type].buttons[i].precalculated_text_position[1], b.b_types[button_type].buttons[i].font_size, rl.White)
		}
	}
	for i := 0; i < len(b.b_types[button_type].buttons); i++ {
		if len(b.b_types[button_type].buttons[i].id) > 6 && b.b_types[button_type].buttons[i].id[len(b.b_types[button_type].buttons[i].id)-7:] == "_select" {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(b.b_types[button_type].buttons[i].position[0]), float32(b.b_types[button_type].buttons[i].position[1]), float32(b.b_types[button_type].normal.Width), float32(b.b_types[button_type].normal.Height))) {
				if rl.IsMouseButtonDown(rl.MouseLeftButton) {
					rl.DrawTexture(*b.b_types[button_type].pressed, b.b_types[button_type].buttons[i].position[0], b.b_types[button_type].buttons[i].position[1], rl.White)
				} else {
					rl.DrawTexture(*b.b_types[button_type].focused, b.b_types[button_type].buttons[i].position[0], b.b_types[button_type].buttons[i].position[1], rl.White)
				}
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !b.pressed_button {
					rl.PlaySound(*b.sound)
					b.b_types[button_type].buttons[i].function(&b.b_types[button_type].buttons[i])
					b.pressed_button = true
					continue
				}
			} else {
				rl.DrawTexture(*b.b_types[button_type].normal, b.b_types[button_type].buttons[i].position[0], b.b_types[button_type].buttons[i].position[1], rl.White)
			}
			rl.DrawText(b.b_types[button_type].buttons[i].text, b.b_types[button_type].buttons[i].precalculated_text_position[0], b.b_types[button_type].buttons[i].precalculated_text_position[1], b.b_types[button_type].buttons[i].font_size, rl.White)
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !b.pressed_button {
		for i := 0; i < len(b.b_types[button_type].buttons); i++ {
			if (len(b.b_types[button_type].buttons[i].id) < 7 || b.b_types[button_type].buttons[i].id[len(b.b_types[button_type].buttons[i].id)-7:] != "_select") && rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(b.b_types[button_type].buttons[i].position[0]), float32(b.b_types[button_type].buttons[i].position[1]), float32(b.b_types[button_type].normal.Width), float32(b.b_types[button_type].normal.Height))) {
				rl.PlaySound(*b.sound)
				b.b_types[button_type].buttons[i].function(&b.b_types[button_type].buttons[i])
				b.pressed_button = true
				break
			}
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for j := 0; j < len(b.b_types[button_type].buttons); j++ {
			if len(b.b_types[button_type].buttons[j].id) > 7 && b.b_types[button_type].buttons[j].id[len(b.b_types[button_type].buttons[j].id)-7:] == "_select" {
				b.Delete(b.b_types[button_type].buttons[j].id)
			}
		}
	}
}

func (b *ButtonType) SetPosition(id string, x, y int32) {
	for i := 0; i < len(b.buttons); i++ {
		if b.buttons[i].id == id {
			b.buttons[i].position[0] = x
			b.buttons[i].position[1] = y
			b.RecalculateTextPosition(i)
		}
	}
}

func (b *ButtonType) SetText(id, text string) {
	for i := 0; i < len(b.buttons); i++ {
		if b.buttons[i].id == id {
			b.buttons[i].text = text
			b.RecalculateTextPosition(i)
		}
	}
}

func (b *ButtonType) GetText(id string) string {
	for i := 0; i < len(b.buttons); i++ {
		if b.buttons[i].id == id {
			return b.buttons[i].text
		}
	}
	return ""
}

func (b *ButtonType) RecalculateTextPosition(i int) {
	b.buttons[i].precalculated_text_position = [2]int32{b.buttons[i].position[0] + b.normal.Width/2 - rl.MeasureText(b.buttons[i].text, b.buttons[i].font_size)/2, b.buttons[i].position[1] + b.normal.Height/2 - b.buttons[i].font_size/2}
}

func (b *ButtonType) DeleteButtonType(id string) {
	for i := 0; i < len(b.buttons); i++ {
		if b.buttons[i].id == id {
			b.buttons = append(b.buttons[:i], b.buttons[i+1:]...)
			i--
		}
	}
}

func (b *Buttons) Delete(id string) {
	for i := 0; i < len(b.b_types); i++ {
		b.b_types[i].DeleteButtonType(id)
	}
}

func (b *Buttons) SetSound(sound *rl.Sound) {
	b.sound = sound
}
