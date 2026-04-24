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
	cursor, err := transactionColl.Find(c, bson.D{{Key: "userId", Value: userId}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível obter as transações do usuário.")
		return
	}
	err = cursor.All(c, &transactionList.List)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível obter as transações do usuário.")
		return
	}

	if transactionList.List == nil {
		transactionList.List = []models.Transaction{}
	}

	c.JSON(http.StatusOK, transactionList)
}

func GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID de transação inválido.")
		return
	}
	userId, _ := c.Get("userId")

	var transaction models.Transaction
	transactionColl := db.GetCollection("transactions")
	err = transactionColl.FindOne(c, bson.D{{Key: "_id", Value: objectId}}).Decode(&transaction)

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
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, "Não foi possível ler os dados da transação.")
		return
	}
	userId, _ := c.Get("userId")

	transaction := models.Transaction{
		UserID:      userId.(bson.ObjectID),
		Type:        newTransaction.Type,
		Value:       newTransaction.Value,
		Description: newTransaction.Description,
		Timestamp:   uint(time.Now().UnixMilli()),
	}

	transactionColl := db.GetCollection("transactions")
	if _, err := transactionColl.InsertOne(c, transaction); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível registrar a transação.")
		return
	}

	c.JSON(http.StatusCreated, "Transação registrada.")
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID de transação inválido.")
		return
	}
	userId, _ := c.Get("userId")

	var transaction models.Transaction
	transactionColl := db.GetCollection("transactions")
	err = transactionColl.FindOne(c, bson.D{{Key: "_id", Value: objectId}}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusNotFound, "Transação não encontrada.")
		return
	}

	if transaction.UserID != userId {
		c.JSON(http.StatusForbidden, "Você não pode deletar uma transação que não é do seu usuário!")
		return
	}

	if _, err := transactionColl.DeleteOne(c, bson.D{{Key: "_id", Value: objectId}}); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível deletar a transação.")
		return
	}
	c.JSON(http.StatusOK, "Transação deletada.")
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID de transação inválido.")
		return
	}
	userId, _ := c.Get("userId")

	var transaction models.Transaction
	transactionColl := db.GetCollection("transactions")
	err = transactionColl.FindOne(c, bson.D{{Key: "_id", Value: objectId}}).Decode(&transaction)

	if err != nil {
		c.JSON(http.StatusNotFound, "Transação não encontrada.")
		return
	}

	if transaction.UserID != userId {
		c.JSON(http.StatusForbidden, "Você não pode editar uma transação que não é do seu usuário!")
		return
	}

	var updatedTransaction models.NewTransaction
	err = c.ShouldBindJSON(&updatedTransaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Não foi possível ler os novos dados da transação.")
		return
	}

	if _, err := transactionColl.UpdateOne(c, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: updatedTransaction}}); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível atualizar a transação.")
		return
	}
	c.JSON(http.StatusOK, "Transação atualizada.")
}

func DeleteAllTransactions(c *gin.Context) {
	userId, _ := c.Get("userId")

	transactionColl := db.GetCollection("transactions")
	if _, err := transactionColl.DeleteMany(c, bson.D{{Key: "userId", Value: userId}}); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível deletar as transações do usuário.")
		return
	}

	c.JSON(http.StatusOK, "Todas as transações do usuário foram deletadas.")
}
