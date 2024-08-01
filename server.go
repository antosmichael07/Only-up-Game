package main

/*import (
	"fmt"
	"os"
	"sync"
	"time"

	lgr "github.com/antosmichael07/Go-Logger"
	tcp "github.com/antosmichael07/Go-TCP-Connection"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func run_server(wg *sync.WaitGroup, wait_for_server *bool, err *error, settings *Settings) {
	server := tcp.NewServer(fmt.Sprintf(":%d", settings.Port))
	server.Logger.Level = lgr.None
	players := map[[64]byte]byte{}
	players_loc := map[[64]byte]rl.Vector2{}

	server.On(event_player_change, func(data *[]byte, conn *tcp.Connection) {
		if len(*data) == 20 {
			players_loc[conn.Token] = rl.Vector2{X: bytes_to_float32((*data)[:4]), Y: bytes_to_float32((*data)[4:8])}

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

	go func() {
		wg.Wait()

		server.Stop()
	}()

	server.OnStart(func() {
		*wait_for_server = false
	})

	*err = server.Start()
}

func saving(file *os.File, players_loc *map[[64]byte]rl.Vector2, saved_highest_loc *rl.Vector2) {
	go func() {
		for {
			time.Sleep(2 * time.Minute)

			if len(*players_loc) > 0 {
				var highest_player rl.Vector2
				for i := range *players_loc {
					if (*players_loc)[i].Y > highest_player.Y {
						highest_player = (*players_loc)[i]
					}
				}

				saved_highest_loc.X = highest_player.X
				saved_highest_loc.Y = highest_player.Y

				_, err := file.WriteAt(append(float32_to_bytes(highest_player.X), float32_to_bytes(highest_player.Y)...), 0)
				if err != nil {
					//logger.Log(lgr.Error, "auto-save: failed to write to save file")
				} else {
					//logger.Log(lgr.Info, "auto-save: saved the highest location")
				}
			}
		}
	}()
}
*/
