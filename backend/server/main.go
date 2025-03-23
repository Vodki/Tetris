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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ws", ws.WebsocketHandler(hub))
	router.GET("/leaderboard", leaderboard.RespondGetLeaderboard())
	router.Run(":8080")
}
