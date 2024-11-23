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
	OBJECT_CARDBOARD_BOX
	OBJECT_PALLETS
	OBJECT_METAL_PIPE
)

const (
	OBJECT_CONTAINER_WIDTH  = 250
	OBJECT_CONTAINER_HEIGHT = 100

	OBJECT_CARDBOARD_BOX_WIDTH  = 25
	OBJECT_CARDBOARD_BOX_HEIGHT = 25

	OBJECT_PALLETS_WIDTH  = 50
	OBJECT_PALLETS_HEIGHT = 10

	OBJECT_METAL_PIPE_WIDTH  = 200
	OBJECT_METAL_PIPE_HEIGHT = 50
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
