package storage

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	log.Println("✅ Connected to MongoDB")
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("webhooks").Collection(collectionName)
}
