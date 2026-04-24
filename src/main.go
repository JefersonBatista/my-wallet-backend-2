package main

import (
	"my-wallet-backend-2/src/db"
	"my-wallet-backend-2/src/routers"
	"my-wallet-backend-2/src/security"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db.Connect()

	engine := gin.Default()
	security.SetCors(engine)
	routers.UseAuthRouter(engine)
	routers.UseTransactionRouter(engine)

	engine.Run()
}
