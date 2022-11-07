package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title    string
	Content  string
	ErrorMsg string
}

func MainPage(c *gin.Context) {
	data := IndexData{"GOAblum", "圖片庫", ""}
	c.HTML(http.StatusOK, "index.html", data)
}

func LoginPage(c *gin.Context) {
	data := new(IndexData)
	data.Title = "Sign in-GOAblum"
	data.Content = "登入"
	data.ErrorMsg = ""
	c.HTML(http.StatusOK, "login.html", data)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	// Check username & password is exist
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Title":    "Sign in-GOAblum",
			"Content":  "登入",
			"ErrorMsg": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Title":    "Sign in-GOAblum",
			"Content":  "登入",
			"ErrorMsg": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	// Check username & password is correct
	if err := Auth(username, password); err == nil {
		/* 將 Token 記錄於 Cookie 中 */
		// c.SetCookie(TOKEN_KEY, TEST_TOKEN, 3600, "/", HOST, false, false)
		// c.HTML(http.StatusOK, "login.html", gin.H{
		// 	"success": "登入成功",
		// })
		MainPage(c)
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Title":    "Sign in-GOAblum",
			"Content":  "登入",
			"ErrorMsg": err,
		})
		return
	}
}

func ginInit() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	//設定靜態資源的讀取
	server.Static("/img", "./img")
	server.StaticFile("/custom", "./template/custom.css")

	server.GET("/", MainPage)
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run(":8888")
}
