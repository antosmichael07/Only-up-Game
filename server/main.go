package main

import (
	"fmt"
	"os"

	lgr "github.com/antosmichael07/Go-Logger"
	tcp "github.com/antosmichael07/Go-TCP-Connection"
)

const (
	event_player_change = iota + 4
	event_player_num
	event_new_player
	event_player_leave
	event_player_kick
	event_i_wanna_leave
	event_side_launcher_launched
)

func main() {
	server := tcp.NewServer(fmt.Sprintf(":%s", os.Args[1]))
	server.Logger.Level = lgr.None
	players := map[[64]byte]byte{}

	server.On(event_player_change, func(data *[]byte, conn *tcp.Connection) {
		server.SendDataToAll(event_player_change, data)
	})

	server.On(event_player_kick, func(data *[]byte, conn *tcp.Connection) {
		server.SendDataToAll(event_player_kick, data)
	})

	server.On(event_i_wanna_leave, func(data *[]byte, conn *tcp.Connection) {
		for i := range server.Connections {
			if server.Connections[i].Token == conn.Token {
				server.Connections[i].ShouldClose = true
				break
			}
		}
	})

	server.On(event_side_launcher_launched, func(data *[]byte, conn *tcp.Connection) {
		server.SendDataToAll(event_side_launcher_launched, data)
	})

	server.OnConnect(func(conn *tcp.Connection) {
		players[conn.Token] = byte(len(server.Connections) - 1)
		server.SendData(conn, event_player_num, &[]byte{byte(len(players) - 1)})
		server.SendDataToAll(event_new_player, &[]byte{byte(len(players))})
	})

	server.OnDisconnect(func(conn *tcp.Connection) {
		server.SendDataToAll(event_player_leave, &[]byte{players[conn.Token]})
		delete(players, conn.Token)
		for i := range players {
			if players[i] > players[conn.Token] {
				players[i]--
			}
		}
	})

	err := server.Start()

	if err != nil {
		fmt.Println(err)
	}
}
