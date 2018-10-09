package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func someGet(c *gin.Context) {
	fmt.Println(c)
}

func init() {
	fmt.Println("Init")
	//添加关闭wraning的提醒
	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/someGet", someGet)
	//http://127.0.0.1:7890/user/444
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") //获取地址栏上的参数
		c.String(http.StatusOK, "Hello %s", name)
	})
	//http://127.0.0.1:7890/user/444/%E9%83%AD%E8%89%B3%E5%B8%85
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	r.Run(":7891") // listen and serve on 0.0.0.0:8080
}
