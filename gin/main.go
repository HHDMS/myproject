package main

import (
	"fmt"
	"gin/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//initDB("tcp://192.168...")
	LoggerMiddleware(initDB)("tcp://192.168...")

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routers.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CheckAuth(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("call checkAuth func", param)
		c.Next()
	}

}

func initDB(connstr string) {
	fmt.Println("初始化数据库", connstr)
}

func LoggerMiddleware(in func(connstr string)) func(connstr string) {
	return func(connstr string) {
		log.Println("call LoggerMiddleware start")
		in(connstr)
		log.Println("call LoggerMiddleware end")
	}
}
