package websocket

import (
	"Tetris/internal/game"
	"Tetris/internal/msg"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins (for development only)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := NewClient(conn)
		hub.AddClient(client)

		msg := msg.WSMessage{
			Type: "id",
			Data: client.id.String(),
		}

		go func() {
			client.Send <- msg
		}()

		eng := game.NewEngine(client.GameSend)
		client.GameRecv = eng.CommandChan

		go eng.Start()
		go client.WritePump()
		go client.readPump()
	}
}
