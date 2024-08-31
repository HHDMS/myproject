package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	//根据模块划分接口
	//根据版本划分接口
	InitCourse(r)
	InitUser(r)
}

func getCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "get course"})
}
