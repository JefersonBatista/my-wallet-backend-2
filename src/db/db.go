package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Database *mongo.Database

func Connect() {
	uri := os.Getenv("MONGO_URI")
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	Database = client.Database("my-wallet")
}

func GetCollection(name string) (coll *mongo.Collection) {
	coll = Database.Collection(name)
	return
}
