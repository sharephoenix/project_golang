package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

case0()

}

//执行返回相关数据
func case0() {
	res, error := execCmd("ls -a")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(res)
}

// 执行 shell 脚本，并输出到文本
func case1()  {
	//定义一个每秒1次输出的shell
	cmdStr := `
#!/bin/bash
for var in {1..10}
do
	sleep 1
    echo "Hello, Welcome ${var} times "
done`
	cmd := exec.Command("bash", "-c",
		cmdStr+" >> shell_output_file.log") //重定向
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

func execCmd(cmdStr string) (res string, err error) {
	args := strings.Split(cmdStr, " ")
	cmd := exec.Command(args[0], args[1:]...)
	resp, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
