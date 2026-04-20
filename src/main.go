package main

import (
	"my-wallet-backend-2/src/controllers"
	"my-wallet-backend-2/src/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/sign-up", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middlewares.Auth, controllers.Logout)

	r.Run()
}
