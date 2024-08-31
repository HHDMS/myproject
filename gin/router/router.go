package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	toc := r.Group("/toc")
	{
		toc.GET("/course", getCourse)
	}

	tob := r.Group("/tob")
	{
		tob.GET("/user", getUser)
	}

}

func getUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user",
	})
}

func getCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get course",
	})
}
