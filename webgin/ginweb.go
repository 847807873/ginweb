package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	//添加默认的字段属性
	/*
	http://127.0.0.1:7891/welcome?lastname=%E5%B0%8F%E7%B1%B3
	*/
	r.GET("/welcome", func(context *gin.Context) {
		//在url没有firstname的时候，展示默认值
		//存在不管有没有值默认值都是不替换的
		//http://127.0.0.1:7891/welcome?lastname=%E5%B0%8F%E7%B1%B3&firstname=
		firstname :=context.DefaultQuery("firstname","郭艳帅123")

		//lastname :=context.Query("lastname")
		//等价的
		lastname := context.Request.URL.Query().Get("lastname")

		context.String(http.StatusOK,"Hello %s %s",firstname,lastname)
	})
	r.Run(":7892") // listen and serve on 0.0.0.0:8080
}
