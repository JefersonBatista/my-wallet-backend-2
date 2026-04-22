package controllers

import (
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTransactions(c *gin.Context) {
	userId, _ := c.Get("userId")

	var user models.User
	userColl := db.GetCollection("users")
	userColl.FindOne(c, bson.D{{Key: "_id", Value: userId}}).Decode(&user)

	var transactionList models.TransactionList
	transactionList.User = user.Name

	transactionColl := db.GetCollection("transactions")
	cursor, _ := transactionColl.Find(c, bson.D{{Key: "userId", Value: userId}})
	cursor.All(c, &transactionList.List)

	c.JSON(http.StatusOK, transactionList)
}
