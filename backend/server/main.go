package main

import (
	ws "Tetris/internal/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	hub := ws.NewHub()
	router := gin.Default()
	router.GET("/ws", ws.WebsocketHandler(hub))
	router.Run(":8080")
}
