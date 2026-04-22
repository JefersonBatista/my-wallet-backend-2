package routers

import (
	"my-wallet-backend-2/src/controllers"
	"my-wallet-backend-2/src/middlewares"

	"github.com/gin-gonic/gin"
)

func UseAuthRouter(engine *gin.Engine) {
	engine.POST("/sign-up", controllers.SignUp)
	engine.POST("/login", controllers.Login)
	engine.POST("/logout", middlewares.Auth, controllers.Logout)
}
