package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/realtime-chat-backend/controller"
	"github.com/kwa0x2/realtime-chat-backend/middlewares"
)

func AuthRoute(router *gin.Engine, authController *controller.AuthController) {
	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.GET("login",authController.GoogleLogin)
		authRoutes.GET("callback", authController.GoogleCallback)
	}
	
}

func UserRoute(router *gin.Engine, userController *controller.UserController) {
	userRoutes := router.Group("/api/v1/user")
	userRoutes.Use(middlewares.JwtMiddleware())
	{
		userRoutes.POST("", userController.InsertUser)
	}
}
