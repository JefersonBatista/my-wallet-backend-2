package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetURI() (uri string) {
	uri = os.Getenv("MONGO_URI")
	return
}

func Connect() (client *mongo.Client, err error) {
	uri := os.Getenv("MONGO_URI")
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err = mongo.Connect(options.Client().ApplyURI(uri))
	return
}
