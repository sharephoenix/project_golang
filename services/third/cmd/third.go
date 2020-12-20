package main

import (
	"encoding/json"
	"example.com/m/services/third/config"
	"example.com/m/services/third/handler"
	"example.com/m/services/third/logic"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var configFile = flag.String("f", "../etc/config.json", "the config file")

func main() {
	flag.Parse()
	fmt.Println("begining!!!", string(*configFile))
	// 配置初始化
	file, _ := os.Open(*configFile)
	fmt.Println("[file]:", file)
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := config.Config{}
	fmt.Println(decoder)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("【Error】", err)
		return
	}

	err = conf.Mysql.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conf.Mysql.DB.Close()
	fmt.Println(conf.Mysql.DbName)

	thirdLogic := logic.TencentLogic{
		&conf,
	}

	tentcentHandler := handler.TencentHandler{
		thirdLogic,
	}
	if conf.EnvMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("[evnMode] release")
	}
	r := gin.Default()
	r.GET("/get/tencentInfo/:appid/:token/:openid", tentcentHandler.GetTencentInfo)
	r.GET("/get/test", tentcentHandler.GetTest)
	defer fmt.Println("finished")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("终于捕获到了panic产生的异常：", err) // 这里的err其实就是panic传入的内容
			fmt.Println("我是defer里的匿名函数，我捕获到panic的异常了，我要recover，恢复过来了。")
		}
	}()
	r.Run(conf.Port)

}
