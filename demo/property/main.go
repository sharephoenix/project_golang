package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string转成int：
	//int, err := strconv.Atoi("1999")
	//fmt.Println(err)
	//fmt.Println(int)

	//string转成int64：
	//int64, err := strconv.ParseInt("99", 10, 64)
	//fmt.Println(err)
	//fmt.Println(int64)
	//int转成string：
	//string := strconv.Itoa(111)
	//fmt.Println(string)
	//int64转成string：
	strin := strconv.FormatInt(99, 10)
	fmt.Println(strin)
}
