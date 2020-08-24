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
	flag.Parse()
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
