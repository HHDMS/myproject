package course

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type course struct {
	Name     string `form:"name" json:"name" binding:"required,alphaunicode"`
	Teacher  string `form:"teacher" json:"teacher" binding:"required,alphaunicode"`
	Duration string `form:"duration" json:"duration" binding:"required,number"`
}

func Add(c *gin.Context) {
	req := &course{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//c.JSON(http.StatusOK, req)
	out := &AddCourseResponse{
		Name:     req.Name,
		Teacher:  req.Teacher,
		Duration: req.Duration,
	}
	c.ProtoBuf(http.StatusOK, out)
}

func Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Update(c *gin.Context) {
	req := &course{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, req)
}

func Delete(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
