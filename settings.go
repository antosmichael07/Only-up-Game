package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/user"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Settings struct {
	Port        uint16
	PlayerLeft  int32
	PlayerRight int32
	PlayerJump  int32
	PlayerKick  int32
}

func create_settings_file() {
	u, _ := user.Current()
	if _, err := os.Stat(fmt.Sprintf("%s/AppData/Local/OnlyUpGame", u.HomeDir)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("%s/AppData/Local/OnlyUpGame", u.HomeDir), 0755)
	}
	if _, err := os.Stat(fmt.Sprintf("%s/AppData/Local/OnlyUpGame/settings.conf", u.HomeDir)); os.IsNotExist(err) {
		os.WriteFile(fmt.Sprintf("%s/AppData/Local/OnlyUpGame/settings.conf", u.HomeDir), default_settings_in_bytes(), 0644)
	}
}

func load_settings() Settings {
	u, _ := user.Current()

	create_settings_file()

	data, err := os.ReadFile(fmt.Sprintf("%s/AppData/Local/OnlyUpGame/settings.conf", u.HomeDir))
	if err != nil {
		return default_settings()
	}

	return Settings{
		Port:        binary.BigEndian.Uint16(data[0:2]),
		PlayerLeft:  int32(binary.BigEndian.Uint32(data[2:6])),
		PlayerRight: int32(binary.BigEndian.Uint32(data[6:10])),
		PlayerJump:  int32(binary.BigEndian.Uint32(data[10:14])),
		PlayerKick:  int32(binary.BigEndian.Uint32(data[14:18])),
	}
}

func save_settings(settings *Settings) {
	u, _ := user.Current()

	create_settings_file()

	data := make([]byte, 18)
	binary.BigEndian.PutUint16(data[0:2], uint16(settings.Port))
	binary.BigEndian.PutUint32(data[2:6], uint32(settings.PlayerLeft))
	binary.BigEndian.PutUint32(data[6:10], uint32(settings.PlayerRight))
	binary.BigEndian.PutUint32(data[10:14], uint32(settings.PlayerJump))
	binary.BigEndian.PutUint32(data[14:18], uint32(settings.PlayerKick))

	os.WriteFile(fmt.Sprintf("%s/AppData/Local/OnlyUpGame/settings.conf", u.HomeDir), data, 0644)
}

func default_settings() Settings {
	return Settings{
		Port:        24680,
		PlayerLeft:  int32(rl.KeyA),
		PlayerRight: int32(rl.KeyD),
		PlayerJump:  int32(rl.KeyW),
		PlayerKick:  int32(rl.KeySpace),
	}
}

func default_settings_in_bytes() []byte {
	return []byte{
		96, 104,
		0, 0, 0, rl.KeyA,
		0, 0, 0, rl.KeyD,
		0, 0, 0, rl.KeyW,
		0, 0, 0, rl.KeySpace,
	}
}

func set_control_setting(control *int32, str_control string, input_box *rl.Texture2D, buttons *Buttons, settings *Settings) {
	key_pressed := rl.GetKeyPressed()
	for key_pressed == 0 {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.DrawTexture(*input_box, int32(rl.GetScreenWidth()/2)-500, 300, rl.White)
		rl.DrawText("PRESS A KEY", int32(rl.GetScreenWidth()/2)-rl.MeasureText("PRESS A KEY", 60)/2, 345, 60, rl.Black)
		rl.EndDrawing()
		key_pressed = rl.GetKeyPressed()
	}

	if key_pressed != rl.KeyEscape {
		*control = key_pressed
		save_settings(settings)

		if key_pressed == rl.KeySpace {
			buttons.b_types[6].SetText(fmt.Sprintf("set-player-%s-setting", str_control), "SPACE")
		} else {
			buttons.b_types[6].SetText(fmt.Sprintf("set-player-%s-setting", str_control), string(*control))
		}
	}
}
