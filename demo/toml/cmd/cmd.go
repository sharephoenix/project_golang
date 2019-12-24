package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Person struct {

	Name string
	Address string
	Age string `toml:"address"`	// 导致 Address 无法读取数据

}

type fconfig struct {
	Owner *Person
	Clients *Person
	Id string
}

func main() {
	var fconfig = fconfig{}
	_, err := toml.DecodeFile("./toml/configs/server.toml", &fconfig)
	if err == nil {
		fmt.Println("read success")
	}
	fmt.Println("read fail")
}