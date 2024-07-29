package main

import (
	"fmt"
	"sync"

	lgr "github.com/antosmichael07/Go-Logger"
	tcp "github.com/antosmichael07/Go-TCP-Connection"
)

func run_server(wg *sync.WaitGroup, wait_for_server *bool, err *error, settings *Settings) {
	server := tcp.NewServer(fmt.Sprintf(":%d", settings.Port))
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

	go func() {
		wg.Wait()

		server.Stop()
	}()

	server.OnStart(func() {
		*wait_for_server = false
	})

	*err = server.Start()
}
