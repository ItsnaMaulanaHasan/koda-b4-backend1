package main

import (
	"gin-practice/middlewares"
	"gin-practice/views"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.AllowPrefic())
	r.Use(middlewares.CorsMiddleware())

	views.SetUpRoutes(r)

	r.Run(":8080")
}
