package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func connect() (client *mongo.Client, err error) {
	uri := os.Getenv("MONGO_URI")
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err = mongo.Connect(options.Client().ApplyURI(uri))
	return
}

func db() (db *mongo.Database) {
	client, _ := connect()
	db = client.Database("my-wallet")
	return
}

func GetCollection(name string) (coll *mongo.Collection) {
	db := db()
	coll = db.Collection(name)
	return
}
