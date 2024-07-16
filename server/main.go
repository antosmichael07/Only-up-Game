package main

import (
	tcp "github.com/antosmichael07/Go-TCP-Connection"
)

const (
	event_player_change = iota + 4
	event_player_num
)

func main() {
	server := tcp.NewServer("192.168.1.127:8080")

	server.On(event_player_change, func(data *[]byte, conn *tcp.Connection) {
		server.SendDataToAll(event_player_change, data)
	})

	server.OnConnect(func(conn *tcp.Connection) {
		server.SendData(conn, event_player_num, &[]byte{byte(len(server.Connections))})
	})

	server.Start()
}
