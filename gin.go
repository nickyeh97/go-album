package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func gin_test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "GOAblum"
	data.Content = "圖片庫"
	c.HTML(http.StatusOK, "index.html", data)
}

func ginWeb() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/", gin_test)
	server.Run(":8888")
}

func ginInit() {
	fmt.Println("123")
}
