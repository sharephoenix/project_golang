package main

import (
	"encoding/json"
	"example.com/m/services/third/handler"
	"example.com/m/services/third/logic"
	"example.com/m/services/user/cmd/api/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var configFile = flag.String("f", "etc/config.json", "the config file")

func main() {
	flag.Parse()
	fmt.Println("begining!!!")
	// 配置初始化
	file, _ := os.Open(*configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := config.Config{}
	fmt.Println(decoder)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("【Error】", err)
	}

	fmt.Println(conf)

	thirdLogic := logic.TencentLogic{
		"",
	}

	tentcentHandler := handler.TencentHandler{
		thirdLogic,
	}

	r := gin.Default()

	r.GET("/get/tencentInfo", tentcentHandler.GetTencentInfo)
}
