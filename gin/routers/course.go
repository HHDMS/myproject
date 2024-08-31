package routers

import (
	"gin/web"
	"github.com/gin-gonic/gin"
)

func InitCourse(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/course", web.GetCourse)
	v1.POST("/course", web.AddCourse)
	//v1.PUT("/course")
}
