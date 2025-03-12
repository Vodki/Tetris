package main

import (
	"Tetris/db"
	"Tetris/db/leaderboard"
	ws "Tetris/internal/websocket"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	hub := ws.NewHub()
	router := gin.Default()
	router.GET("/ws", ws.WebsocketHandler(hub))
	router.GET("/scores", func(c *gin.Context) {
		scores, err := leaderboard.GetScores()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch scores"})
			return
		}
		c.JSON(200, scores)
	})
	router.Run(":8080")
}
