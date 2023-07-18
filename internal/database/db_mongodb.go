package database

import (
	"context"
	"knowledge-api/internal/config"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect_MongoDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.ConnetStringMongoDB)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	time.AfterFunc(10*time.Second, func() {
		client.Disconnect(context.Background())
	})

	db := client.Database(os.Getenv("MONGO_DATABASE"))

	return db, nil
}
