package websocket

import "sync"

type Hub struct {
	clients map[*Client]bool
	mutex   sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) AddClient(client *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.clients[client] = true
}

func (h *Hub) RemoveClient(client *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	delete(h.clients, client)
}
