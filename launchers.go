package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SideLauncher struct {
	Rect           rl.Rectangle
	Power          float32
	AnimationTimer float32
}

func (side_launcher *SideLauncher) Update(side_launcher_textures *[2][4]rl.Texture2D) {
	if side_launcher.AnimationTimer > 0 {
		side_launcher.AnimationTimer -= rl.GetFrameTime()
	}

	if side_launcher.Power < 0 {
		if side_launcher.AnimationTimer > 1 {
			rl.DrawTexture(side_launcher_textures[1][3], int32(side_launcher.Rect.X)-10, int32(side_launcher.Rect.Y), rl.White)
		} else if side_launcher.AnimationTimer > .5 {
			rl.DrawTexture(side_launcher_textures[1][2], int32(side_launcher.Rect.X)-10, int32(side_launcher.Rect.Y), rl.White)
		} else if side_launcher.AnimationTimer > 0 {
			rl.DrawTexture(side_launcher_textures[1][1], int32(side_launcher.Rect.X)-10, int32(side_launcher.Rect.Y), rl.White)
		} else {
			rl.DrawTexture(side_launcher_textures[1][0], int32(side_launcher.Rect.X)-10, int32(side_launcher.Rect.Y), rl.White)
		}
	} else if side_launcher.Power >= 0 {
		if side_launcher.AnimationTimer > 1 {
			rl.DrawTexture(side_launcher_textures[0][3], int32(side_launcher.Rect.X), int32(side_launcher.Rect.Y), rl.White)
		} else if side_launcher.AnimationTimer > .5 {
			rl.DrawTexture(side_launcher_textures[0][2], int32(side_launcher.Rect.X), int32(side_launcher.Rect.Y), rl.White)
		} else if side_launcher.AnimationTimer > 0 {
			rl.DrawTexture(side_launcher_textures[0][1], int32(side_launcher.Rect.X), int32(side_launcher.Rect.Y), rl.White)
		} else {
			rl.DrawTexture(side_launcher_textures[0][0], int32(side_launcher.Rect.X), int32(side_launcher.Rect.Y), rl.White)
		}
	}
}
