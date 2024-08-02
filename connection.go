package main

import (
	"sync"
	"time"

	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	event_player_change = iota + 4
	event_player_num
	event_new_player
	event_player_leave
	event_player_kick
	event_i_wanna_leave
	event_side_launcher_launched
	event_launcher_launched
)

func connection(players *[]Player, wg *sync.WaitGroup, player_num *byte, remove_player *byte, should_close_connection *bool, wg_disconnect *sync.WaitGroup, client *tcp.Client, wait_player_num_wg *sync.WaitGroup, side_launchers *[]SideLauncher, player_loc *rl.Vector2, launchers *[]Launcher) {
	wait_player_num := byte(255)

	client.On(event_player_leave, func(data *[]byte) {
		*remove_player = (*data)[0]
	})

	client.On(event_new_player, func(data *[]byte) {
		if *player_num == 255 {
			for i := 0; i < int((*data)[0]); i++ {
				*players = append(*players, NewPlayer())
			}
			wait_player_num_wg.Wait()
			*player_num = wait_player_num
			(*players)[*player_num].Position = *player_loc
			wg.Done()
		} else {
			*players = append(*players, NewPlayer())
		}
	})

	client.On(event_player_num, func(data *[]byte) {
		wait_player_num = (*data)[0]
		player_loc.X = bytes_to_float32((*data)[1:5])
		player_loc.Y = bytes_to_float32((*data)[5:])

		wait_player_num_wg.Done()
	})

	client.On(event_player_change, func(data *[]byte) {
		if (*data)[15] == *player_num {
			return
		}

		x := bytes_to_float32((*data)[:4])
		y := bytes_to_float32((*data)[4:8])
		(*players)[(*data)[15]].Gravity = bytes_to_float32((*data)[8:12])
		(*players)[(*data)[15]].Keys[0] = (*data)[12]
		(*players)[(*data)[15]].Keys[1] = (*data)[13]
		(*players)[(*data)[15]].Keys[2] = (*data)[14]
		(*players)[(*data)[15]].SideLauncherPower = bytes_to_float32((*data)[16:])

		if (*players)[(*data)[15]].Position.X+50 < x || (*players)[(*data)[15]].Position.X-50 > x {
			(*players)[(*data)[15]].Position.X = x
		}

		if (*players)[(*data)[15]].Position.Y+50 < y || (*players)[(*data)[15]].Position.Y-50 > y {
			(*players)[(*data)[15]].Position.Y = y
		}
	})

	client.On(event_player_kick, func(data *[]byte) {
		(*players)[(*data)[0]].SideLauncherPower = bytes_to_float32((*data)[1:5])
		(*players)[(*data)[5]].Kicking = true
	})

	client.On(event_side_launcher_launched, func(data *[]byte) {
		if (*side_launchers)[(*data)[0]].AnimationTimer <= 0 {
			(*side_launchers)[(*data)[0]].AnimationTimer = 2
		}
	})

	client.On(event_launcher_launched, func(data *[]byte) {
		if (*launchers)[(*data)[0]].AnimationTimer <= 0 {
			(*launchers)[(*data)[0]].AnimationTimer = 2
		}
	})

	go client.Listen()

	for *player_num == 255 {
		time.Sleep(200 * time.Millisecond)

		if *should_close_connection {
			client.SendData(event_i_wanna_leave, &[]byte{})

			client.Disconnect()
			wg_disconnect.Done()
			return
		}
	}

	data_sending(client, players, player_num, should_close_connection, wg_disconnect)
}

func data_sending(client *tcp.Client, players *[]Player, player_num *byte, should_close_connection *bool, wg_disconnect *sync.WaitGroup) {
	last_position := rl.NewVector2(0, 0)
	last_direction := int8(0)
	last_gravity := float32(0)
	last_input := [3]byte{0, 0, 0}
	last_side_launcher_power := float32(0)

	for !*should_close_connection {
		if last_position.X != (*players)[*player_num].Position.X || last_position.Y != (*players)[*player_num].Position.Y || last_direction != (*players)[*player_num].Direction || last_gravity != (*players)[*player_num].Gravity || last_input[0] != (*players)[*player_num].Keys[0] || last_input[1] != (*players)[*player_num].Keys[1] || last_input[2] != (*players)[*player_num].Keys[2] || last_side_launcher_power != (*players)[*player_num].SideLauncherPower {
			send_data(client, players, player_num)
		}

		last_position = (*players)[*player_num].Position
		last_direction = (*players)[*player_num].Direction
		last_gravity = (*players)[*player_num].Gravity
		last_input = (*players)[*player_num].Keys
		last_side_launcher_power = (*players)[*player_num].SideLauncherPower

		time.Sleep(20 * time.Millisecond)
	}

	client.SendData(event_i_wanna_leave, &[]byte{})

	client.Disconnect()
	wg_disconnect.Done()
}

func send_data(client *tcp.Client, players *[]Player, player_num *byte) {
	data := make([]byte, 20)
	x := float32_to_bytes((*players)[*player_num].Position.X)
	y := float32_to_bytes((*players)[*player_num].Position.Y)
	gravity := float32_to_bytes((*players)[*player_num].Gravity)
	side_launcher_power := float32_to_bytes((*players)[*player_num].SideLauncherPower)

	copy(data[:4], x)
	copy(data[4:8], y)
	copy(data[8:12], gravity)
	data[12] = (*players)[*player_num].Keys[0]
	data[13] = (*players)[*player_num].Keys[1]
	data[14] = (*players)[*player_num].Keys[2]
	data[15] = *player_num
	copy(data[16:], side_launcher_power)

	client.SendData(event_player_change, &data)
}
