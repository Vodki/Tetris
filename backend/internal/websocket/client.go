package websocket

import (
	"encoding/json"
	"log"

	"Tetris/internal/msg"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	GameSend chan msg.GameMessage
	Send     chan msg.WSMessage
	id       uuid.UUID
	GameRecv chan msg.WSMessage
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:     conn,
		GameSend: make(chan msg.GameMessage),
		Send:     make(chan msg.WSMessage),
		id:       uuid.New(),
	}
}

func (c *Client) WritePump() {
	defer c.conn.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			toSend, err := json.Marshal(message)
			if err != nil {
				log.Print("Error marshaling JSON:", err)
			}
			c.conn.WriteMessage(websocket.TextMessage, toSend)
		case message, ok := <-c.GameSend:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			toSend, err := json.Marshal(message)
			if err != nil {
				log.Print("Error marshaling JSON:", err)
			}
			c.conn.WriteMessage(websocket.TextMessage, toSend)
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		var msg msg.WSMessage
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Printf("Error decoding JSON: %v", err)
		}

		log.Printf("Received message: %+v", msg) // Print the decoded message
		c.handleMessage(msg)                     // Handle the message
	}
}

func (c *Client) handleMessage(msg msg.WSMessage) {
	msgType := string(msg.Type)
	switch msgType {
	case "game":
		c.GameRecv <- msg
	case "start":
		c.GameRecv <- msg
	}
}
