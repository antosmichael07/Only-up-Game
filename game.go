package main

import (
	"errors"
	"sync"
	"time"

	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func init_game() ([]Player, []rl.Rectangle, []SideLauncher, []Launcher, rl.Camera2D) {
	players := []Player{}

	collision_rects := []rl.Rectangle{
		rl.NewRectangle(25, 160, 250, 25),
		rl.NewRectangle(25, 150, 80, 10),
		rl.NewRectangle(140, 150, 25, 10),
	}

	side_launchers := []SideLauncher{
		NewSideLauncher(25, 90, 15, &collision_rects),
	}

	launchers := []Launcher{}

	camera := rl.NewCamera2D(rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)), rl.NewVector2(225, 0), 0, 4)

	return players, collision_rects, side_launchers, launchers, camera
}

func game_loop(should_close_connection *bool, client *tcp.Client, player_textures *[][3]rl.Texture2D, arrow *rl.Texture2D, buttons *Buttons, is_game_menu_open *bool, side_launcher_textures *[2][4]rl.Texture2D, err *error, settings *Settings) {
	players, collision_rects, side_launchers, launchers, camera := init_game()
	player_num := byte(255)
	remove_player := byte(255)
	just_closed_game_menu := false

	var wg_disconnect sync.WaitGroup
	wg_disconnect.Add(2)

	var wait_player_num_wg sync.WaitGroup
	wait_player_num_wg.Add(1)

	var wg sync.WaitGroup
	wg.Add(1)
	go connection(&players, &wg, &player_num, &remove_player, should_close_connection, &wg_disconnect, client, &wait_player_num_wg)
	go func() {
		time.Sleep(5 * time.Second)
		if player_num == 255 {
			*err = errors.New("starter data wasn't received correctly, please try again")
			*should_close_connection = true
			wg.Done()
		}
	}()
	wg.Wait()

	for !*should_close_connection {
		window_manager()
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode2D(camera)
		update_camera(&players, &camera, &player_num)

		for i := 0; i < len(players); i++ {
			players[i].Update(&collision_rects, &side_launchers, &launchers, player_textures, &players)
		}
		players[player_num].DrawArrow(arrow)

		for i := 0; i < len(collision_rects); i++ {
			rl.DrawRectangleRec(collision_rects[i], rl.Black)
		}
		for i := 0; i < len(side_launchers); i++ {
			side_launchers[i].Update(side_launcher_textures)
		}
		for i := 0; i < len(launchers); i++ {
			rl.DrawRectangleRec(launchers[i].Rect, rl.Green)
		}

		rl.EndMode2D()
		if *is_game_menu_open {
			rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight())+300, rl.Fade(rl.Black, 0.6))
			buttons.Draw(5)
			if rl.IsKeyPressed(rl.KeyEscape) && !just_closed_game_menu {
				*is_game_menu_open = false
				just_closed_game_menu = true
			}
		} else {
			players[player_num].Input(settings)
			players[player_num].Kick(&players, &player_num, client, settings)
		}

		if rl.IsKeyPressed(rl.KeyEscape) {
			*is_game_menu_open = true
		}
		just_closed_game_menu = false

		rl.EndDrawing()

		if remove_player != 255 {
			players = append(players[:remove_player], players[remove_player+1:]...)
			if player_num > remove_player {
				player_num--
			}
			remove_player = 255
		}
	}
	wg_disconnect.Done()

	wg_disconnect.Wait()
}

func update_camera(players *[]Player, camera *rl.Camera2D, player_num *byte) {
	camera.Target.X = rl.Lerp(camera.Target.X, (*players)[*player_num].Position.X+12.5, 0.05*(*players)[*player_num].FrameTime)
	camera.Target.Y = rl.Lerp(camera.Target.Y, (*players)[*player_num].Position.Y, 0.05*(*players)[*player_num].FrameTime)
}
