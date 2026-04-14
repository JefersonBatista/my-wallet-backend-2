package main

import (
	"my-wallet-backend-2/src/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", getUsers)

	r.GET("/db-uri", getDBURI)

	r.Run()
}

type User struct {
	Name string `json:"name" bson:"name"`
}

func getUsers(c *gin.Context) {
	client, _ := db.Connect()
	// println(client)
	coll := client.Database("my-wallet").Collection("users")
	// println(coll)

	filter := bson.D{{}}

	var users []User
	cursor, err := coll.Find(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cursor.All(c, &users)

	c.JSON(http.StatusOK, users)
}

func getDBURI(c *gin.Context) {
	uri := db.GetURI()
	c.JSON(http.StatusOK, gin.H{
		"uri": uri,
	})
}
