package main

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_Main(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", Ping)
	v1 := router.Group("/v1")
	{
		v1.GET("/ping/:jump", PingV1)
		v1.GET("/ping", PingV1)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/ping", PingV2)
	}

	router.GET("login/", Login)
	router.Run(":80")
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
