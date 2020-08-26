package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var configFile = flag.String("f", "config.json", "hahah")

func main() {

}

func demo2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
func demo0() {
	//list := new([]int)
	//list = append(list, 1)
	//fmt.Println(list)
}

func demo1() {
	in := 9
	defer fmt.Println("11111", in)
	flag.Parse()
	in = 10
	defer fmt.Println("22222", in)
	fmt.Println(*configFile)
	cmd := exec.Command("pwd")
	byt, err := cmd.Output()
	if err == nil {
		fmt.Println("result", string(byt))
	} else {
		fmt.Println("error::", err.Error())
	}
	os.Chdir(string(byt))

	file, e := os.Open(*configFile)
	if e != nil {
		fmt.Println("error:", e.Error())
		return
	}
	decoder := json.NewDecoder(file)
	var r map[string]string
	decoder.Decode(r)
	fmt.Println(r)
}
