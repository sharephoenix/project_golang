package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"os"
	"project_golang/common/baseresponse"
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

	// 全局中间件 middleware-校验 Authorization token 是否合法
	r.Use(func(context *gin.Context) {

		fmt.Println("Im a dummy!11111111", conf)

		// Pass on to the next-in-chain

		context.Next()

	}, func (context *gin.Context) {

		fmt.Println("Im a dummy!22222222", conf)

		// Pass on to the next-in-chain

		context.Next()

	})

	/*局部中间件*/
	// 获取当前用户登录信息
	r.GET("/logininfo", func(context *gin.Context) {
		jwtToken := context.Request.Header["Authorization"]
		token := context.Request.Header["Token"]
		if jwtToken == nil || len(jwtToken) <= 0{
			resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"头部缺少 Authorization"})
			context.JSON(200, resp)
			context.Abort()
			return
		}
		if len(token) <= 0 {
			resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"头部缺少 token"})
			context.JSON(200, resp)
			context.Abort()
			return
		}
		jwt, err := logic2.BackGenToken(jwtToken[0], conf.Auth.AccessSecret)
		if err != nil {
			resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{err.Error()})
			context.JSON(200, resp)
			context.Abort()
			return
		}
		if jwt["usr"] != token[0] {
			resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"Authorization 失效"})
			context.JSON(200, resp)
			context.Abort()
		}
		context.Next()
	} ,userHandler.GetUser)
	r.GET("/sendCode/:mobile", userHandler.SendCode)
	r.GET("/getCode/:mobile", userHandler.GetCode)
	r.POST("/register", func(context *gin.Context) {
		context.Next()
	},userHandler.Register(conf.Auth.AccessSecret))

	r.Run(conf.Port)
}




