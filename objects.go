package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Object struct {
	Position rl.Vector2
	Texture  *rl.Texture2D
}

type PreObject struct {
	Texture       rl.Texture2D
	CollisionRect CollisionRect
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
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 250, Height: 100}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 25, Height: 25}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 50, Height: 10}, false}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 200, Height: 50}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 20, Height: 150}, false}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 150, Height: 20}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 20, Height: 10}, true}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 0, Width: 4, Height: 9}, false}},
	{rl.Texture2D{}, CollisionRect{rl.Rectangle{X: 0, Y: 2, Width: 20, Height: 28}, false}},
}

func NewObject(x, y float32, obj uint, collision_rects *[]CollisionRect) Object {
	*collision_rects = append(*collision_rects, NewCollisionRect(x + pre_objects[obj].CollisionRect.Rect.X, -y + pre_objects[obj].CollisionRect.Rect.Y, pre_objects[obj].CollisionRect.Rect.Width, pre_objects[obj].CollisionRect.Rect.Height, pre_objects[obj].CollisionRect.Climbable))

	return Object{
		Position: rl.NewVector2(x, -y),
		Texture:  &pre_objects[obj].Texture,
	}
}

func (o *Object) Draw() {
	rl.DrawTexture(*o.Texture, int32(o.Position.X), int32(o.Position.Y), rl.White)
}
