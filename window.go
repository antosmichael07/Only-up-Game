package main

import rl "github.com/gen2brain/raylib-go/raylib"

var refresh_rate int32

func init_window() {
	rl.SetTraceLogLevel(rl.LogError)
	monitor := rl.GetCurrentMonitor()
	rl.InitWindow(int32(rl.GetMonitorWidth(monitor)), int32(rl.GetMonitorHeight(monitor)), "Only up!")
	refresh_rate = int32(rl.GetMonitorRefreshRate(monitor))
	rl.SetTargetFPS(refresh_rate)
	rl.ToggleFullscreen()
	rl.SetExitKey(-1)
}

func fps_reducer() {
	if rl.IsWindowFocused() {
		rl.SetTargetFPS(refresh_rate)
	} else {
		rl.SetTargetFPS(30)
	}
}
