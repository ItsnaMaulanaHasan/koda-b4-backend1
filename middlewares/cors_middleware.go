package middlewares

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "google.com")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "content-type")
		ctx.Next()
	}
}
