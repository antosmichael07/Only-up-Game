package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Position rl.Vector2
	Width    float32
	Height   float32
	Texture  *rl.Texture2D
}

type PreObject struct {
	Width         float32
	Height        float32
	Texture       rl.Texture2D
	CollisionRect rl.Rectangle
}

const (
	OBJECT_CONTAINER_RED = iota
	OBJECT_CONTAINER_GREEN
	OBJECT_CONTAINER_BLUE
	OBJECT_CONTAINER_YELLOW
	OBJECT_CARDBOARD_BOX
	OBJECT_PALLETS
	OBJECT_METAL_PIPE
	OBJECT_METAL_SUPPORT
	OBJECT_METAL_SUPPORT_HORIZONTAL
	OBJECT_BRICK
	OBJECT_BOLT
	OBJECT_TOILET
)

var pre_objects = []PreObject{
	{250, 100, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}},
	{250, 100, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}},
	{250, 100, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}},
	{250, 100, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}},
	{25, 25, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 25, Height: 25}},
	{50, 10, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 50, Height: 10}},
	{200, 50, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 200, Height: 50}},
	{20, 150, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 20, Height: 150}},
	{150, 20, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 150, Height: 20}},
	{20, 10, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 20, Height: 10}},
	{4, 9, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 0, Width: 4, Height: 9}},
	{20, 30, rl.Texture2D{}, rl.Rectangle{X: 0, Y: 2, Width: 20, Height: 28}},
}

func NewObject(x, y float32, obj uint, collision_rects *[]rl.Rectangle) Object {
	*collision_rects = append(*collision_rects, rl.NewRectangle(x+pre_objects[obj].CollisionRect.X, -y+pre_objects[obj].CollisionRect.Y, pre_objects[obj].CollisionRect.Width, pre_objects[obj].CollisionRect.Height))

	return Object{
		Position: rl.NewVector2(x, -y),
		Width:    pre_objects[obj].Width,
		Height:   pre_objects[obj].Height,
		Texture:  &pre_objects[obj].Texture,
	}
}

func (o *Object) Draw() {
	rl.DrawTexture(*o.Texture, int32(o.Position.X), int32(o.Position.Y), rl.White)
}

func NewCollisionRect(x, y, width, height float32) rl.Rectangle {
	return rl.Rectangle{
		X:      x,
		Y:      -y,
		Width:  width,
		Height: height,
	}
}
