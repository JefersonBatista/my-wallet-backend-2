package routers

import (
	"my-wallet-backend-2/src/controllers"
	"my-wallet-backend-2/src/middlewares"

	"github.com/gin-gonic/gin"
)

func UseTransactionRouter(engine *gin.Engine) {
	routeGroup := engine.Group("/transactions", middlewares.Auth)
	routeGroup.GET("", controllers.GetTransactions)
	routeGroup.GET("/:id", controllers.GetTransactionById)
	routeGroup.POST("", controllers.RegisterTransaction)
	routeGroup.DELETE("/all", controllers.DeleteAllTransactions)
	routeGroup.DELETE("/:id", controllers.DeleteTransaction)
	routeGroup.PUT("/:id", controllers.UpdateTransaction)
}
