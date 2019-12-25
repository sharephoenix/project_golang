package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"os"
	"project_golang/services/user/cmd/api/config"
	"project_golang/services/user/handler"
	logic2 "project_golang/services/user/logic"
	"project_golang/services/user/model"
)

var configFile = flag.String("f", "etc/config.json", "the config file")

func main() {
	//flag.Parse()
	//var c config.Config
	//conf.MustLoad(*configFile, &c)

	// 配置初始化
	file, _ := os.Open("etc/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := config.Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}


	biz := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := biz.Ping().Result()
	fmt.Println(pong, err)

	// 业务初始化
	userModel := model.UserModel{
		biz,
	}

	logic := logic2.UserLogic{
		userModel,
	}

	// 路由
	userHandler := handler.UserHandler{
		logic,
	}
	r := gin.Default()
	r.GET("/user/:mobile", userHandler.GetUser)
	r.POST("/register",userHandler.Register)
	r.Run(conf.Port)
}
