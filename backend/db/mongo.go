package db

import (
    "context"
    "fmt"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get MongoDB URI from environment variables
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        return fmt.Errorf("MONGODB_URI environment variable not set")
    }

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return fmt.Errorf("failed to connect to MongoDB: %v", err)
    }

    // Ping the database to verify the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to ping MongoDB: %v", err)
    }

    // Get database name from environment variables (fallback to "leaderboard_dev")
    dbName := os.Getenv("MONGODB_DB")
    if dbName == "" {
        dbName = "leaderboard_dev"
    }

    DB = client.Database(dbName)
    Client = client
    fmt.Println("Connected to MongoDB!")
    return nil
}

func Disconnect() {
    if Client == nil {
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _ = Client.Disconnect(ctx)
    fmt.Println("Disconnected from MongoDB.")
}