package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SideLauncher struct {
	Rect           rl.Rectangle
	Power          float32
	AnimationTimer float32
}

func NewSideLauncher(x, y, power float32, collision_rects *[]CollisionRect) SideLauncher {
	if power < 0 {
		*collision_rects = append(*collision_rects, NewCollisionRect(x+36, -y+10, 14, 40, true))
		return SideLauncher{rl.NewRectangle(x+10, -y, 40, 60), power, 0}
	} else {
		*collision_rects = append(*collision_rects, NewCollisionRect(x, -y+10, 14, 40, true))
		return SideLauncher{rl.NewRectangle(x, -y, 40, 60), power, 0}
	}
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

type Launcher struct {
	Rect  rl.Rectangle
	Power float32
}

func NewLauncher(x, y, power float32, collision_rects *[]CollisionRect) Launcher {
	*collision_rects = append(*collision_rects, NewCollisionRect(x, -y, 60, 10, false))
	return Launcher{rl.NewRectangle(x+5, -y-1, 50, 11), -power}
}

func (launcher *Launcher) Update(launcher_texture *rl.Texture2D) {
	rl.DrawTexture(*launcher_texture, int32(launcher.Rect.X)-5, int32(launcher.Rect.Y)+1, rl.White)
}
