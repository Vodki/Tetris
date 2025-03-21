package leaderboard

import (
	"Tetris/db"
	"context"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type leaderboardEntry struct {
	Username string `bson:"username" json:"username"`
	Score    int    `bson:"score" json:"score"`
}

func NewLeaderboardEntry(username string, score int) *leaderboardEntry {
	return &leaderboardEntry{
		Username: username,
		Score:    score,
	}
}

type leaderboardMessage struct {
	Type string             `json:"type"`
	Data []leaderboardEntry `json:"data"`
}

func GetLeaderboard() ([]leaderboardEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("LeaderboardEntry")

	var entries []leaderboardEntry
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &entries); err != nil {
		return nil, err
	}

	return entries, err
}

func AddScore(entry leaderboardEntry) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("LeaderboardEntry")

	var entries []leaderboardEntry
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, &entries); err != nil {
		return err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Score < entries[j].Score
	})

	if entry.Score > entries[0].Score {
		filter := bson.D{{Key: "username", Value: entries[0].Username}, {Key: "score", Value: entries[0].Score}}
		if _, err := collection.ReplaceOne(ctx, filter, entry); err != nil {
			return err
		}
	}
	return nil
}

func RespondGetLeaderboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		leaderboard, err := GetLeaderboard()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, leaderboard)
	}
}
