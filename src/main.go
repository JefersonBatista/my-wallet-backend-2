package main

import (
	"my-wallet-backend-2/src/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	engine := gin.Default()
	routers.UseAuthRouter(engine)
	routers.UseTransactionRouter(engine)

	engine.Run()
}
