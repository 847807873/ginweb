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
		id :=context.Query("id")
		page :=context.DefaultQuery("page","0")
		name :=context.PostForm("name")
		message :=context.PostForm("message")

		context.String(http.StatusOK,"id %s ,page %s , name %s , message %s ",id,page,name,message)

	})

	r.Run(":7895")
}

