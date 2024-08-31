package routers

import (
	"gin/middleware"
	"gin/web"
	"github.com/gin-gonic/gin"
)

func InitUser(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middleware.Auth())
	v1.GET("/user", web.GetUser)
	v1.POST("/user", web.Adduser)
	//v1.PUT("/user")
}
