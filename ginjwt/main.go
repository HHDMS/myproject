package main

import (
	"ginjwt/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	routers.InitRouters(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
