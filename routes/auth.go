package routes

import (
	"gin-practice/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, authController *controllers.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.PATCH("/forgot-password/:id", authController.ForgotPassword)
	}
}
