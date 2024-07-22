package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var refresh_rate int32

func init_window() {
	rl.SetTraceLogLevel(rl.LogError)
	monitor := rl.GetCurrentMonitor()
	rl.InitWindow(int32(rl.GetMonitorWidth(monitor)), int32(rl.GetMonitorHeight(monitor)), "Only up!")
	refresh_rate = int32(rl.GetMonitorRefreshRate(monitor))
	rl.SetTargetFPS(refresh_rate)
	rl.ToggleFullscreen()
	rl.SetExitKey(-1)

	icon := rl.LoadImage("resources/textures/player_10.png")
	rl.SetWindowIcon(*icon)
	rl.UnloadImage(icon)
}

func window_manager() {
	if rl.IsWindowFocused() {
		if !rl.IsWindowFullscreen() {
			rl.ToggleFullscreen()
		}
		rl.SetTargetFPS(refresh_rate)
	} else {
		if rl.IsWindowFullscreen() {
			rl.ToggleFullscreen()
		}
		rl.SetTargetFPS(30)
	}

	if rl.WindowShouldClose() {
		os.Exit(0)
	}
}
