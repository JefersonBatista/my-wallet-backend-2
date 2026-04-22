package routers

import (
	"my-wallet-backend-2/src/controllers"
	"my-wallet-backend-2/src/middlewares"

	"github.com/gin-gonic/gin"
)

func UseTransactionRouter(engine *gin.Engine) {
	engine.GET("/transactions", middlewares.Auth, controllers.GetTransactions)
}
