package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (player *Player) Input() {
	if rl.IsKeyDown(rl.KeyA) {
		player.Keys[0] = 1
	} else {
		player.Keys[0] = 0
	}

	if rl.IsKeyDown(rl.KeyD) {
		player.Keys[1] = 1
	} else {
		player.Keys[1] = 0
	}

	if rl.IsKeyDown(rl.KeyW) {
		player.Keys[2] = 1
	} else {
		player.Keys[2] = 0
	}
}
