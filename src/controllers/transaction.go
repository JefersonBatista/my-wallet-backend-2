package controllers

import (
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/models"
	"net/http"
	"time"

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

	if transactionList.List == nil {
		transactionList.List = []models.Transaction{}
	}

	c.JSON(http.StatusOK, transactionList)
}

func GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := bson.ObjectIDFromHex(id)
	userId, _ := c.Get("userId")

	var transaction models.Transaction
	transactionColl := db.GetCollection("transactions")
	err := transactionColl.FindOne(c, bson.D{{Key: "_id", Value: objectId}}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusNotFound, "Transação não encontrada.")
		return
	}

	if transaction.UserID != userId {
		c.JSON(http.StatusForbidden, "Você não pode obter uma transação que não é do seu usuário!")
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func RegisterTransaction(c *gin.Context) {
	var newTransaction models.NewTransaction
	c.ShouldBindJSON(&newTransaction)
	userId, _ := c.Get("userId")

	transaction := models.Transaction{
		UserID:      userId.(bson.ObjectID),
		Type:        newTransaction.Type,
		Value:       newTransaction.Value,
		Description: newTransaction.Description,
		Timestamp:   uint(time.Now().UnixMilli()),
	}

	transactionColl := db.GetCollection("transactions")
	transactionColl.InsertOne(c, transaction)

	c.JSON(http.StatusCreated, "Transação registrada.")
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := bson.ObjectIDFromHex(id)
	userId, _ := c.Get("userId")

	var transaction models.Transaction
	transactionColl := db.GetCollection("transactions")
	err := transactionColl.FindOne(c, bson.D{{Key: "_id", Value: objectId}}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusNotFound, "Transação não encontrada.")
		return
	}

	if transaction.UserID != userId {
		c.JSON(http.StatusForbidden, "Você não pode deletar uma transação que não é do seu usuário!")
		return
	}

	transactionColl.DeleteOne(c, bson.D{{Key: "_id", Value: objectId}})
	c.JSON(http.StatusOK, "Transação deletada.")
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := bson.ObjectIDFromHex(id)
	userId, _ := c.Get("userId")

	var transaction models.Transaction
	transactionColl := db.GetCollection("transactions")
	err := transactionColl.FindOne(c, bson.D{{Key: "_id", Value: objectId}}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusNotFound, "Transação não encontrada.")
		return
	}

	if transaction.UserID != userId {
		c.JSON(http.StatusForbidden, "Você não pode editar uma transação que não é do seu usuário!")
		return
	}

	var updatedTransaction models.NewTransaction
	c.ShouldBindJSON(&updatedTransaction)

	transactionColl.UpdateOne(c, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: updatedTransaction}})
	c.JSON(http.StatusOK, "Transação atualizada.")
}

func DeleteAllTransactions(c *gin.Context) {
	userId, _ := c.Get("userId")

	transactionColl := db.GetCollection("transactions")
	transactionColl.DeleteMany(c, bson.D{{Key: "userId", Value: userId}})

	c.JSON(http.StatusOK, "Todas as transações do usuário foram deletadas.")
}
