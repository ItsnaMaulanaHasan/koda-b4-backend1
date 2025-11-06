package routes

import (
	"gin-practice/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {

	userController := controllers.NewUserController()
	authController := controllers.NewAuthController(userController)

	users := r.Group("/users")
	{
		users.GET("", userController.GetAllUser)
		users.GET("/:id", userController.GetUserById)
		users.POST("", userController.CreateUser)
		users.PATCH("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.PATCH("/forgot-password/:id", authController.ForgotPassword)
	}

}
