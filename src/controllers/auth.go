package controllers

import (
	"fmt"
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var newUser models.NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, "Não foi possível ler os dados do usuário.")
		return
	}

	userColl := db.GetCollection("users")
	findSameEmail := userColl.FindOne(c, bson.D{{Key: "email", Value: newUser.Email}})
	if conflict := findSameEmail.Decode(&models.User{}) == nil; conflict {
		c.JSON(http.StatusConflict, "Um usuário com esse email já está cadastrado.")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível criar o usuário.")
		return
	}

	var user models.User
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.PasswordHash = string(passwordHash)
	if _, err := userColl.InsertOne(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível criar o usuário.")
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("Usuário %s cadastrado.", user.Name))
}

func Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, "Não foi possível ler os dados de login.")
		return
	}

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
	if _, err := sessionColl.InsertOne(c, models.Session{
		UserID: user.ID,
		Token:  token,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível criar a sessão.")
		return
	}

	c.JSON(http.StatusOK, token)
}

func Logout(c *gin.Context) {
	token := c.GetString("token")

	sessionColl := db.GetCollection("sessions")
	if _, err := sessionColl.DeleteOne(c, bson.D{{Key: "token", Value: token}}); err != nil {
		c.JSON(http.StatusInternalServerError, "Não foi possível encerrar a sessão.")
		return
	}
	c.JSON(http.StatusOK, "Sessão encerrada.")
}
