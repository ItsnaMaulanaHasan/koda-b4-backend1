package views

import (
	"gin-practice/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/users", controllers.DataUsers.GetAllUser)
	r.GET("/users/:id", controllers.DataUsers.GetUserById)
	r.POST("/users", controllers.DataUsers.CreateUser)
	r.PATCH("/users/:id", controllers.DataUsers.UpdateUser)
	r.DELETE("/users/:id", controllers.DataUsers.DeleteUser)
}
