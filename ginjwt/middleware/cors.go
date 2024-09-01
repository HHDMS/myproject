package middleware

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Called Cors1")
	}
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{
			"Origin", "Content-Type", "Content-Length",
		},
		AllowMethods: []string{
			"GET", "POST", "DELETE", "PUT", "OPTIONS",
		},
	})
}
