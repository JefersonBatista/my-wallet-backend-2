package controllers

import (
	"fmt"
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var newUser models.NewUser
	c.ShouldBindJSON(&newUser)

	userColl := db.GetCollection("users")
	findSameEmail := userColl.FindOne(c, bson.D{{Key: "email", Value: newUser.Email}})
	if conflict := findSameEmail.Decode(&models.User{}) == nil; conflict {
		c.JSON(http.StatusConflict, "Um usuário com esse email já está cadastrado.")
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	var user models.User
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.PasswordHash = string(passwordHash)
	userColl.InsertOne(c, user)

	c.JSON(http.StatusCreated, fmt.Sprintf("Usuário %s cadastrado.", user.Name))
}

func Login(c *gin.Context) {
	var login models.Login
	c.ShouldBindJSON(&login)

	userColl := db.GetCollection("users")
	var user models.User
	err := userColl.FindOne(c, bson.D{{Key: "email", Value: login.Email}}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, "Nenhum usuário com esse email está cadastrado.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Senha incorreta!")
		return
	}

	token := uuid.NewString()
	sessionColl := db.GetCollection("sessions")
	sessionColl.InsertOne(c, models.Session{
		UserID: user.ID,
		Token:  token,
	})

	c.JSON(http.StatusOK, token)
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	if token == "" {
		c.JSON(http.StatusUnauthorized, "Você não está autorizado!")
		return
	}

	var session models.Session
	sessionColl := db.GetCollection("sessions")
	err := sessionColl.FindOne(c, bson.D{{Key: "token", Value: token}}).Decode(&session)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Você não está autorizado!")
		return
	}

	var user models.User
	userColl := db.GetCollection("users")
	userColl.FindOne(c, bson.D{{Key: "_id", Value: session.UserID}}).Decode(&user)
	sessionColl.DeleteOne(c, bson.D{{Key: "token", Value: token}})
	c.JSON(http.StatusOK, fmt.Sprintf("Sessão do usuário %s encerrada.", user.Name))
}
