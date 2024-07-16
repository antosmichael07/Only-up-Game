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
)

func connection(player *Player, player_2 *Player, wg *sync.WaitGroup) {
	client := tcp.NewClient("192.168.1.127:8080")
	client.Connect()
	player_num := byte(255)

	client.On(event_player_num, func(data *[]byte) {
		player_num = (*data)[0]
		wg.Done()
	})

	client.On(event_player_change, func(data *[]byte) {
		if (*data)[15] == player_num {
			return
		}

		x := bytes_to_float32((*data)[:4])
		y := bytes_to_float32((*data)[4:8])
		player_2.Gravity = bytes_to_float32((*data)[8:12])
		player_2.Keys[0] = (*data)[12]
		player_2.Keys[1] = (*data)[13]
		player_2.Keys[2] = (*data)[14]

		if player_2.Position.X+15 < x || player_2.Position.X-15 > x {
			player_2.Position.X = x
		}

		if player_2.Position.Y+15 < y || player_2.Position.Y-15 > y {
			player_2.Position.Y = y
		}
	})

	go client.Listen()

	for player_num == 255 {
		time.Sleep(200 * time.Millisecond)
	}

	data_sending(&client, player, &player_num)
}

func data_sending(client *tcp.Client, player *Player, player_num *byte) {
	last_position := rl.NewVector2(0, 0)
	last_direction := int8(0)
	last_gravity := float32(0)
	last_input := [3]byte{0, 0, 0}

	for !rl.WindowShouldClose() {
		if last_position.X != player.Position.X || last_position.Y != player.Position.Y || last_direction != player.Direction || last_gravity != player.Gravity || last_input[0] != player.Keys[0] || last_input[1] != player.Keys[1] || last_input[2] != player.Keys[2] {
			send_data(client, player, player_num)
		}

		last_position = player.Position
		last_direction = player.Direction
		last_gravity = player.Gravity
		last_input = player.Keys

		time.Sleep(20 * time.Millisecond)
	}
}

func send_data(client *tcp.Client, player *Player, player_num *byte) {
	data := make([]byte, 16)
	x := float32_to_bytes(player.Position.X)
	y := float32_to_bytes(player.Position.Y)
	gravity := float32_to_bytes(player.Gravity)

	copy(data[:4], x)
	copy(data[4:8], y)
	copy(data[8:12], gravity)
	data[12] = player.Keys[0]
	data[13] = player.Keys[1]
	data[14] = player.Keys[2]
	data[15] = *player_num

	client.SendData(event_player_change, &data)
}
