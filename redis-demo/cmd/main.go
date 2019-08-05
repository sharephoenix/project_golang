package main

import (
	"fmt"
	"kratos-master/pkg/cache/redis"
)

func init()  {      //init 用于初始化一些参数，先于main执行

}


func main() {
	fmt.Println("this is redis")

	conn,err := redis.Dial("tcp","10.1.210.69:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	defer conn.Close()


	res,err := conn.Do("HSET","student","name","jack")
	fmt.Println(res,err)
	res1,err := redis.String(conn.Do("HGET","student","name"))
	fmt.Printf("res:%s,error:%v",res1,err)
}