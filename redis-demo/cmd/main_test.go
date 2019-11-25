package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

/// CURL 请求
//curl --header "A-BBb aaaa" -H "A-CcV:BBBBBBB" --data "account=alexluan&password=qwe123" localhost:8009/login-post
//{"account":"alexluan","pwd":"qwe123"}

func Test_Main(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", Ping)
	router.GET("/ping302", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/ping/:jump", PingV1)
		v1.GET("/ping", PingV1)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/ping", PingV2)
	}

	router.GET("/login", Login)
	router.POST("/login-post", LoginPost)
	router.Run(":8009")
}


func Login(c *gin.Context) {
	account := c.Query("account")
	pwd := c.Query("pwd")
	c.JSON(200, gin.H{
		"account": account,
		"pwd": pwd,
	})
	//account := c.PostForm("account")

}

func LoginPost(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("password")
	for k,v :=range c.Request.Header {
		fmt.Println(k,v)
	}
	c.JSON(200, gin.H{
		"account": account,
		"pwd": pwd,
	})
}

/************/
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PingV1(c *gin.Context) {
	jump := c.Param("jump")
	var mm map[string]interface{} = make(map[string]interface{})
	mm["aa"] = "bbbb"
	mm["message"] = "pong"
	mm["jump"] = jump
	c.JSON(200, mm)
}

func PingV2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong v2",
	})
}
