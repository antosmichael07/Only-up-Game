package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"

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

type Vector2 struct {
	X float32
	Y float32
}

func main() {
	var server tcp.Server

	logger, lgr_err := lgr.NewLogger("SERVER", "logs", true)
	if lgr_err != nil {
		logger.Output.File = false
		logger.Log(lgr.Error, "failed to open logger files, logging to console only")
	}

	if len(os.Args) < 2 {
		logger.Log(lgr.Warning, "usage: server.exe <port>")
		logger.Log(lgr.Info, "starting server on default port 24680...")
		server = tcp.NewServer(":24680")
	} else {
		server = tcp.NewServer(fmt.Sprintf(":%s", os.Args[1]))
	}

	server.Logger.Level = lgr.None
	players := map[[64]byte]byte{}
	players_loc := map[[64]byte]Vector2{}

	server.On(event_player_change, func(data *[]byte, conn *tcp.Connection) {
		if len(*data) == 20 {
			players_loc[conn.Token] = Vector2{X: bytes_to_float32((*data)[:4]), Y: bytes_to_float32((*data)[4:8])}

			server.SendDataToAll(event_player_change, data)
		}
	})

	server.On(event_player_kick, func(data *[]byte, conn *tcp.Connection) {
		if len(*data) == 6 {
			server.SendDataToAll(event_player_kick, data)
		}
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
		if len(*data) == 1 {
			server.SendDataToAll(event_side_launcher_launched, data)
		}
	})

	server.OnConnect(func(conn *tcp.Connection) {
		players[conn.Token] = byte(len(server.Connections) - 1)
		server.SendData(conn, event_player_num, &[]byte{byte(len(players) - 1)})
		server.SendDataToAll(event_new_player, &[]byte{byte(len(players))})
	})

	server.OnDisconnect(func(conn *tcp.Connection) {
		server.SendDataToAll(event_player_leave, &[]byte{players[conn.Token]})
		delete(players_loc, conn.Token)
		delete(players, conn.Token)
		for i := range players {
			if players[i] > players[conn.Token] {
				players[i]--
			}
		}
	})

	logger.Log(lgr.Info, "server is starting...")

	server.OnStart(func() {
		logger.Log(lgr.Info, "server started")
	})

	go func() {
		skip := false
		for {
			str := ""
			fmt.Scanf("%s", &str)

			if skip {
				skip = false
				continue
			}

			if str == "" {
				continue
			}

			time_now := time.Now().String()[:19]
			if logger.Output.File {
				if logger.OpenedFile != time_now[:10] {
					logger.File.Close()
					if file, err := os.OpenFile(fmt.Sprintf("./%s/%s.txt", logger.Directory, time_now[:10]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
						logger.File = file
						logger.OpenedFile = time_now[:10]
					}
				}

				logger.File.WriteString(fmt.Sprintf("[%s] %s\n", time_now, str))
			}

			switch str {
			case "stop":
				logger.Log(lgr.Info, "stopping the server...")
				server.Stop()
				return

			default:
				logger.Log(lgr.Error, "unknown command")
			}

			skip = true
		}
	}()

	err := server.Start()

	if err != nil {
		fmt.Println(err)
	}
}

func bytes_to_float32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}
