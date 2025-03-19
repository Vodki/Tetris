package leaderboard

import (
	"Tetris/db"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type leaderboardEntry struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name" json:"name"`
	Score int    `bson:"score" json:"score"`
}

func GetScores() ([]leaderboardEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("LeaderboardEntry")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var scores []leaderboardEntry
	if err = cursor.All(ctx, &scores); err != nil {
		return nil, err
	}

	return scores, nil
}

func GetLeaderboard() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("LeaderboardEntry")

	var entries []bson.M
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var entry bson.M
		if err := cursor.Decode(&entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)

	}
	return entries, err

}

func AddScore(entry leaderboardEntry) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("leaderboard")
	var result bson.M
	if err := collection.FindOne(ctx, bson.D{{Key: "rank", Value: 10}}).Decode(&result); err != nil {
		fmt.Fprintf(os.Stderr, "Fetching lowest score failed: %v", err)
		return nil, err
	}

	if result["score"].(int) > entry.Score {
		return nil, nil
	}

	return collection.InsertOne(ctx, entry)
}
