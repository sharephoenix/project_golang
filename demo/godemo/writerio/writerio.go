package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//a := "Hello, playground"
	//fmt.Println([]byte(a))

	a := "Hello, playground"
	buf := new(bytes.Buffer)
	buf.ReadFrom(strings.NewReader(a))
	fmt.Println(buf.Bytes())
}
