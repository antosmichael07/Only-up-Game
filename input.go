package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (player *Player) Input(settings *Settings) {
	if rl.IsKeyDown(settings.PlayerLeft) {
		player.Keys[0] = 1
	} else {
		player.Keys[0] = 0
	}

	if rl.IsKeyDown(settings.PlayerRight) {
		player.Keys[1] = 1
	} else {
		player.Keys[1] = 0
	}

	if rl.IsKeyDown(settings.PlayerJump) {
		player.Keys[2] = 1
	} else {
		player.Keys[2] = 0
	}
}
