package middlewares

import (
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Você não está autorizado!")
		return
	}

	var session models.Session
	sessionColl := db.GetCollection("sessions")
	err := sessionColl.FindOne(c, bson.D{{Key: "token", Value: token}}).Decode(&session)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Você não está autorizado!")
		return
	}

	c.Set("userId", session.UserID)
	c.Set("token", token)
}
