package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	log.Println("Executing middleware")

	c.Next()

	log.Println("Middleware execution complete")
}
