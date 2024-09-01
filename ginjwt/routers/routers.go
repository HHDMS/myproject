package routers

import (
	"ginjwt/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.Cors(), middleware.Auth())
	//课程相关接口
	initCourse(api)
	//用户相关接口
	initUser(api)

	notAuthApi := r.Group("/api")
	notAuthApi.Use(middleware.Cors())
	//登录接口
	initLogin(notAuthApi)
}
