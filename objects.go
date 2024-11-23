package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Position rl.Vector2
	Width    float32
	Height   float32
	Texture  *rl.Texture2D
}

const (
	OBJECT_CONTAINER = iota
)

func NewObject(x, y, width, height float32, texture *rl.Texture2D, collision_rects *[]rl.Rectangle) Object {
	*collision_rects = append(*collision_rects, rl.NewRectangle(x, y, width, height))

	return Object{
		Position: rl.NewVector2(x, y),
		Width:    width,
		Height:   height,
		Texture:  texture,
	}
}

func (o *Object) Draw() {
	rl.DrawTexture(*o.Texture, int32(o.Position.X), int32(o.Position.Y), rl.White)
}
