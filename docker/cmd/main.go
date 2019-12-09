package main

import (
	"fmt"
	"time"
)

type UserBase struct {
	Name *string `json:"name, omitempty"`
	Sex *int `json:"sex"`
} 

type User struct{
	Name *string `json:"name"`
	Address string `json:"address"`
	Item UserBase `json:"item"`
}

func main() {
	//var user User
	////user.Name = "alexluan"
	//user.Address = "shanghai"
	//a := "eric"
	//user.Item.Name = &a
	//
	//by, err := json.Marshal(user)
	//if err == nil {
	//	fmt.Println(user.Name)
	//	fmt.Println(string(by))
	//}
	fmt.Println("before")
	now := time.Now().UnixNano() / 1e6
	fmt.Println("start")
	fmt.Println(now)
	fmt.Println("end")

	var names []string

	names = append(names, "alex")
	names = append(names, []string{"alex", "jump"}...)
	fmt.Println(names)
	}