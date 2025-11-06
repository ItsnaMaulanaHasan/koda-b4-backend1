package routes

import (
	"gin-practice/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userController *controllers.UserController) {
	users := r.Group("/users")
	{
		users.GET("", userController.GetAllUser)
		users.GET("/:id", userController.GetUserById)
		users.POST("", userController.CreateUser)
		users.PATCH("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
		users.PATCH("/:id/upload-profile", userController.UploadProfile)
	}
}
