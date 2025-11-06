package routes

import (
	"gin-practice/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()
	authController := controllers.NewAuthController(userController)

	SetupUserRoutes(r, userController)
	SetupAuthRoutes(r, authController)
}
