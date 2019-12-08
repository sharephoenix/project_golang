package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type UserBase struct {
	Name string `json:"name"`
}

type LessBase struct {
	Age string `json:"age"`
}

type User struct {
	UserBase
	LessBase
	Name string `json:"name"`
	Address string `json:"address"`
}

func main() {

	var user User
	user.Name = "alexluan"
	user.Address = "shanghai"
	jbyt, _ := json.Marshal(user)

	fmt.Println(string(jbyt))
	//fmt.Println("start")
	//defer fmt.Println("end")
	//TestSyncMap()

	//maproutine.TestSyncMap()

	//maproutine.TestMap()
}

func TestSyncMap() {
		var m sync.Map
		m.Store("method", "eth_getBlockByHash")
		m.Store("jsonrpc", "2.0")
		//value, ok := m.Load("method")
		//t.Logf("value=%v,ok=%v\n", value, ok)
		//f := func(key, value string) {
		//
		//}
		//f("method", "eth_getBlockByHash")
		m.Range(func(key, value interface{}) bool {
			fmt.Println("range k:%v,v=%v\n", key, value)
			return true
		})
}