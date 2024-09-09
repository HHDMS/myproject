package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"limiter-breaker/breaker"
	"limiter-breaker/limiter"
	"limiter-breaker/middleware"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", middleware.Limiter(limiter.NewLimiter(3*time.Second, 4)), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	b := breaker.NewBreaker(4, 4, 2, time.Second*15)
	r.GET("/ping1", func(c *gin.Context) {
		err := b.Exec(func() error {
			value, _ := c.GetQuery("value")
			if value == "a" {
				return errors.New("value 为 a 返回错误")
			}
			return nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong1",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
