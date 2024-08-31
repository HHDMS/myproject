package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	//id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

type User struct {
	Name  string `form:"name" binding:"required"`
	Phone string `form:"phone" binding:"required,e164"`
}

func Adduser(ctx *gin.Context) {
	req := &User{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, req)
}
