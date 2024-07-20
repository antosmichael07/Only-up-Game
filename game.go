package main

import (
	"sync"

	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func init_game() ([]Player, []rl.Rectangle, []SideLauncher, []Launcher, rl.Camera2D) {
	players := []Player{}

	collision_rects := []rl.Rectangle{
		rl.NewRectangle(25, 125, 250, 25),
	}

	side_launchers := []SideLauncher{
		{rl.NewRectangle(25, 100, 25, 25), 8},
		{rl.NewRectangle(250, 100, 25, 25), -8},
	}

	launchers := []Launcher{}

	camera := rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)), rl.NewVector2(225, 0), 0, 4)

	return players, collision_rects, side_launchers, launchers, camera
}

func game_loop(should_close_connection *bool, client *tcp.Client, player_textures *[][3]rl.Texture2D, arrow *rl.Texture2D) {
	players, collision_rects, jumpers, side_launchers, camera := init_game()
	player_num := byte(255)
	remove_player := byte(255)

	var wg_disconnect sync.WaitGroup
	wg_disconnect.Add(2)

	var wait_player_num_wg sync.WaitGroup
	wait_player_num_wg.Add(1)

	var wg sync.WaitGroup
	wg.Add(1)
	go connection(&players, &wg, &player_num, &remove_player, should_close_connection, &wg_disconnect, client, &wait_player_num_wg)
	wg.Wait()

	for !*should_close_connection {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(camera)
		update_camera(&players, &camera, &player_num)

		players[player_num].Input()
		for i := 0; i < len(players); i++ {
			players[i].Update(&collision_rects, &jumpers, &side_launchers, player_textures, &players)
		}
		players[player_num].Kick(&players, &player_num, client)
		players[player_num].DrawArrow(arrow)

		for i := 0; i < len(collision_rects); i++ {
			rl.DrawRectangleRec(collision_rects[i], rl.Black)
		}
		for i := 0; i < len(jumpers); i++ {
			rl.DrawRectangleRec(jumpers[i].Rect, rl.Red)
		}
		for i := 0; i < len(side_launchers); i++ {
			rl.DrawRectangleRec(side_launchers[i].Rect, rl.Green)
		}

		rl.EndMode2D()
		rl.EndDrawing()

		if remove_player != 255 {
			players = append(players[:remove_player], players[remove_player+1:]...)
			if player_num > remove_player {
				player_num--
			}
			remove_player = 255
		}

		if rl.IsKeyPressed(rl.KeyEscape) {
			*should_close_connection = true
		}
	}
	wg_disconnect.Done()

	wg_disconnect.Wait()
}

func update_camera(players *[]Player, camera *rl.Camera2D, player_num *byte) {
	camera.Target.Y = rl.Lerp(camera.Target.Y, (*players)[*player_num].Position.Y, 0.05*(*players)[*player_num].FrameTime)
}
