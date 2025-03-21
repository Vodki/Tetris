package main

import (
	"Tetris/db"
	"Tetris/db/leaderboard"
	ws "Tetris/internal/websocket"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	hub := ws.NewHub()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your React app URL
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	router.GET("/ws", ws.WebsocketHandler(hub))
	router.GET("/leaderboard", leaderboard.RespondGetLeaderboard())
	router.Run(":8080")
}
