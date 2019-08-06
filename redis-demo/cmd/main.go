package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type User struct {
	Name string	`json:"name"`
	Address string	`json:"address"`
	Sex int64	`json:"sex"`
}

func main() {
	fmt.Println("this is redis")
	defer fmt.Println("THIS REDIS END")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	Set(client)
	Get(client)
}

func Set(client *redis.Client) {
	user := User{
		"alexluan",
		"anhui",
		1,
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	err0 := client.Set("user", string(jsonBytes), 0).Err()
	if err0 != nil {
		panic(err)
	}

	val, err := client.Get("user").Result()
	if err == nil {
		var userr User
		e := json.Unmarshal([]byte(val), &userr)
		if e == nil {
			fmt.Println("struct", userr)
		} else {
			fmt.Println("errrrror", e.Error())
		}
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("user", val)
}

func Get(client *redis.Client) {
	val2, err := client.Get("key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val2)
	}
}