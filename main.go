package main

import (
	"gin-practice/views"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	views.SetUpRoutes(r)

	r.Run(":8080")
}
