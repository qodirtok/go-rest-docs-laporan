package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/helper"
)

func LoggerMiddleware(c *gin.Context) {
	log.Println("Executing middleware")

	c.Next()

	log.Println("Middleware execution complete")
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		context.Next()
	}
}
