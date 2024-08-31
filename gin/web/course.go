package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	//id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

type Course struct {
	Name    string `form:"name" binding:"required"`
	Teacher string `form:"teacher" binding:"required"`
	Price   string `form:"price" binding:"number"`
}

func AddCourse(ctx *gin.Context) {
	req := &Course{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, req)
}
