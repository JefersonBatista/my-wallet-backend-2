package main

import (
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/models"
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

	r.GET("/ping", pingPong)
	r.GET("/users", getUsers)

	r.Run()
}

func pingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getUsers(c *gin.Context) {
	coll := db.GetCollection("users")

	filter := bson.D{{}}

	var users []models.User
	cursor, err := coll.Find(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cursor.All(c, &users)

	c.JSON(http.StatusOK, users)
}
