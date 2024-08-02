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
	event_launcher_launched
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
	saved_highest_loc := Vector2{X: 100, Y: 100}
	auto_save := true

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

	server.On(event_launcher_launched, func(data *[]byte, conn *tcp.Connection) {
		if len(*data) == 1 {
			server.SendDataToAll(event_launcher_launched, data)
		}
	})

	server.OnConnect(func(conn *tcp.Connection) {
		players[conn.Token] = byte(len(server.Connections) - 1)
		to_send := []byte{byte(len(players) - 1)}
		to_send = append(to_send, float32_to_bytes(saved_highest_loc.X)...)
		to_send = append(to_send, float32_to_bytes(saved_highest_loc.Y)...)
		server.SendData(conn, event_player_num, &to_send)
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

	var file *os.File
	if _, err := os.Stat("save"); os.IsNotExist(err) {
		logger.Log(lgr.Warning, "no save file found, creating a new one")
		if file, err = os.Create("save"); err != nil {
			logger.Log(lgr.Error, "failed to create save file")
		} else {
			_, err = file.Write(append(float32_to_bytes(100), float32_to_bytes(100)...))
			if err != nil {
				logger.Log(lgr.Error, "failed to write to save file")
			}

			go saving(file, &players_loc, &logger, &saved_highest_loc, &auto_save)
		}
	} else {
		if file, err = os.OpenFile("save", os.O_RDWR, 0644); err != nil {
			logger.Log(lgr.Error, "failed to open save file")
		} else {
			data := make([]byte, 8)
			_, err = file.ReadAt(data, 0)
			if err != nil {
				logger.Log(lgr.Error, "failed to read save file")
			} else {
				saved_highest_loc.X = bytes_to_float32(data[:4])
				saved_highest_loc.Y = bytes_to_float32(data[4:])
			}

			go saving(file, &players_loc, &logger, &saved_highest_loc, &auto_save)
		}
	}

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
				if len(players_loc) > 0 {
					var highest_player Vector2
					for i := range players_loc {
						if (players_loc)[i].Y > highest_player.Y {
							highest_player = (players_loc)[i]
						}
					}

					saved_highest_loc.X = highest_player.X
					saved_highest_loc.Y = highest_player.Y

					_, err := file.WriteAt(append(float32_to_bytes(highest_player.X), float32_to_bytes(highest_player.Y)...), 0)
					if err != nil {
						logger.Log(lgr.Error, "auto-save: failed to write to save file")
						logger.Log(lgr.Info, "write 'Y' to stop the server without saving the highest location")

						dump := ""
						fmt.Scanln(&dump)

						var input string
						fmt.Scanf("%s", &input)

						if input != "Y" {
							logger.Log(lgr.Info, "canceled stopping the server")
							continue
						}
					} else {
						logger.Log(lgr.Info, "auto-save: saved the highest location")
					}
				} else {
					logger.Log(lgr.Warning, "no players online to save the highest location")
					logger.Log(lgr.Info, "write 'Y' to stop the server without saving the highest location")

					dump := ""
					fmt.Scanln(&dump)

					var input string
					fmt.Scanf("%s", &input)

					if input != "Y" {
						logger.Log(lgr.Info, "canceled stopping the server")
						continue
					}
				}

				logger.Log(lgr.Info, "stopping the server...")
				server.Stop()
				return

			case "save":
				if len(players_loc) > 0 {
					var highest_player Vector2
					for i := range players_loc {
						if (players_loc)[i].Y > highest_player.Y {
							highest_player = (players_loc)[i]
						}
					}

					saved_highest_loc.X = highest_player.X
					saved_highest_loc.Y = highest_player.Y

					_, err := file.WriteAt(append(float32_to_bytes(highest_player.X), float32_to_bytes(highest_player.Y)...), 0)
					if err != nil {
						logger.Log(lgr.Error, "failed to write to save file")
					} else {
						logger.Log(lgr.Info, "saved the highest location")
					}
				} else {
					logger.Log(lgr.Warning, "no players online to save the highest location")
				}

			case "listen":
				server.ShouldListen = !server.ShouldListen
				if server.ShouldListen {
					logger.Log(lgr.Info, "server is now listening for new connections")
				} else {
					logger.Log(lgr.Info, "server is not listening for new connections")
				}

			case "auto-save":
				auto_save = !auto_save
				if auto_save {
					logger.Log(lgr.Info, "auto-save is now enabled")
				} else {
					logger.Log(lgr.Info, "auto-save is now disabled")
				}

			default:
				logger.Log(lgr.Error, "unknown command")
			}

			skip = true
		}
	}()

	err := server.Start()

	if err != nil {
		logger.Log(lgr.Error, "%s", err)
	}
}

func bytes_to_float32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func float32_to_bytes(f float32) []byte {
	bytes := make([]byte, 4)
	bits := math.Float32bits(f)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func saving(file *os.File, players_loc *map[[64]byte]Vector2, logger *lgr.Logger, saved_highest_loc *Vector2, auto_save *bool) {
	go func() {
		for {
			time.Sleep(2 * time.Minute)

			if *auto_save && len(*players_loc) > 0 {
				var highest_player Vector2
				for i := range *players_loc {
					if (*players_loc)[i].Y > highest_player.Y {
						highest_player = (*players_loc)[i]
					}
				}

				saved_highest_loc.X = highest_player.X
				saved_highest_loc.Y = highest_player.Y

				_, err := file.WriteAt(append(float32_to_bytes(highest_player.X), float32_to_bytes(highest_player.Y)...), 0)
				if err != nil {
					logger.Log(lgr.Error, "auto-save: failed to write to save file")
				} else {
					logger.Log(lgr.Info, "auto-save: saved the highest location")
				}
			}
		}
	}()
}
