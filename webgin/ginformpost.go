package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form_post", func(context *gin.Context) {
		//通过key 不是通过body
		message := context.PostForm("message")

		nick := context.DefaultPostForm("nick", "默认值")
		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/post", func(context *gin.Context) {
		id :=context.Query("id") //通过url获取
		page :=context.DefaultQuery("page","0") //通过url获取
		name :=context.PostForm("name") //通过 from-data类型设置获取
		message :=context.PostForm("message")

		context.String(http.StatusOK,"id %s ,page %s , name %s , message %s ",id,page,name,message)

	})

	r.POST("/postmap", func(context *gin.Context) {
		//请求的url
		//http://127.0.0.1:7895/postmap?names[first]=thinkerou&names[second]=tianou&ids[a]=1234&ids[b]=hello
		ids :=context.QueryMap("ids")
		names :=context.QueryMap("names")
		context.String(http.StatusOK,"ids: %v; names: %v", ids, names)
	})

	r.Run(":7895")
}

