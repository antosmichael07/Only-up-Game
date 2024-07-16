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
)

func connection(players *[]Player, wg *sync.WaitGroup, player_num *byte, remove_player *byte, should_close_connection *bool, wg_disconnect *sync.WaitGroup, client *tcp.Client) {
	wait_player_num := byte(255)

	client.On(event_player_leave, func(data *[]byte) {
		*remove_player = (*data)[0]
	})

	client.On(event_new_player, func(data *[]byte) {
		if *player_num == 255 {
			for i := 0; i < int((*data)[0]); i++ {
				*players = append(*players, NewPlayer())
			}
			*player_num = wait_player_num
			wg.Done()
		} else {
			*players = append(*players, NewPlayer())
		}
	})

	client.On(event_player_num, func(data *[]byte) {
		wait_player_num = (*data)[0]
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

		if (*players)[(*data)[15]].Position.X+50 < x || (*players)[(*data)[15]].Position.X-50 > x {
			(*players)[(*data)[15]].Position.X = x
		}

		if (*players)[(*data)[15]].Position.Y+50 < y || (*players)[(*data)[15]].Position.Y-50 > y {
			(*players)[(*data)[15]].Position.Y = y
		}
	})

	client.On(event_player_kick, func(data *[]byte) {
		(*players)[(*data)[0]].SideLauncherPower = bytes_to_float32((*data)[1:])
	})

	go client.Listen()

	for *player_num == 255 {
		time.Sleep(200 * time.Millisecond)
	}

	data_sending(client, players, player_num, should_close_connection, wg_disconnect)
}

func data_sending(client *tcp.Client, players *[]Player, player_num *byte, should_close_connection *bool, wg_disconnect *sync.WaitGroup) {
	last_position := rl.NewVector2(0, 0)
	last_direction := int8(0)
	last_gravity := float32(0)
	last_input := [3]byte{0, 0, 0}

	for !rl.WindowShouldClose() && !*should_close_connection {
		if last_position.X != (*players)[*player_num].Position.X || last_position.Y != (*players)[*player_num].Position.Y || last_direction != (*players)[*player_num].Direction || last_gravity != (*players)[*player_num].Gravity || last_input[0] != (*players)[*player_num].Keys[0] || last_input[1] != (*players)[*player_num].Keys[1] || last_input[2] != (*players)[*player_num].Keys[2] {
			send_data(client, players, player_num)
		}

		last_position = (*players)[*player_num].Position
		last_direction = (*players)[*player_num].Direction
		last_gravity = (*players)[*player_num].Gravity
		last_input = (*players)[*player_num].Keys

		time.Sleep(20 * time.Millisecond)
	}

	client.Disconnect()
	wg_disconnect.Done()
}

func send_data(client *tcp.Client, players *[]Player, player_num *byte) {
	data := make([]byte, 16)
	x := float32_to_bytes((*players)[*player_num].Position.X)
	y := float32_to_bytes((*players)[*player_num].Position.Y)
	gravity := float32_to_bytes((*players)[*player_num].Gravity)

	copy(data[:4], x)
	copy(data[4:8], y)
	copy(data[8:12], gravity)
	data[12] = (*players)[*player_num].Keys[0]
	data[13] = (*players)[*player_num].Keys[1]
	data[14] = (*players)[*player_num].Keys[2]
	data[15] = *player_num

	client.SendData(event_player_change, &data)
}
