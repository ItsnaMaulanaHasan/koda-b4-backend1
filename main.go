// @title           API Documentation
// @version         1.0
// @description     Dokumentasi REST API menggunakan Gin dan Swagger

// @host      localhost:8080
// @BasePath  /

package main

import (
	"gin-practice/middlewares"
	"gin-practice/routes"

	_ "gin-practice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.AllowPrefic())
	r.Use(middlewares.CorsMiddleware())

	routes.SetUpRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
