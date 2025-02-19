package websocket

import "github.com/gorilla/websocket"

func broadcast(hub *Hub, message []byte) {
	hub.mutex.RLock()
	defer hub.mutex.RUnlock()

	for client := range hub.clients {
		// Use a goroutine to avoid blocking the hub
		go func(c *websocket.Conn) {
			err := c.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				// Handle error (e.g., remove connection)
				hub.RemoveClient(client)
				c.Close()
			}
		}(client.conn)
	}
}
