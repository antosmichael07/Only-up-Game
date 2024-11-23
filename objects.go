package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Position rl.Vector2
	Width    float32
	Height   float32
	Texture  *rl.Texture2D
}

type PreObject struct {
	Width   float32
	Height  float32
	Texture rl.Texture2D
}

const (
	OBJECT_CONTAINER = iota
	OBJECT_CARDBOARD_BOX
	OBJECT_PALLETS
	OBJECT_METAL_PIPE
	OBJECT_METAL_SUPPORT
	OBJECT_METAL_SUPPORT_HORIZONTAL
)

var pre_objects = []PreObject{
	{250, 100, rl.Texture2D{}},
	{25, 25, rl.Texture2D{}},
	{50, 10, rl.Texture2D{}},
	{200, 50, rl.Texture2D{}},
	{25, 150, rl.Texture2D{}},
	{150, 25, rl.Texture2D{}},
}

func NewObject(x, y float32, obj uint, collision_rects *[]rl.Rectangle) Object {
	*collision_rects = append(*collision_rects, rl.NewRectangle(x, y, pre_objects[obj].Width, pre_objects[obj].Height))

	return Object{
		Position: rl.NewVector2(x, y),
		Width:    pre_objects[obj].Width,
		Height:   pre_objects[obj].Height,
		Texture:  &pre_objects[obj].Texture,
	}
}

func (o *Object) Draw() {
	rl.DrawTexture(*o.Texture, int32(o.Position.X), int32(o.Position.Y), rl.White)
}
