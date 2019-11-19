package logic

import (
	"fmt"
	"testing"
)

func Test_Register_test(t *testing.T) {
	err := Register("18817322819", "qwe123", "alexlaun",
		"shanghai", 0)
	if err == nil {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败")
	}
}

func Test_CheckRegister(t *testing.T) {
	account := "18817322819"
	res, err := CheckRegister(account)
	if err == nil {
		fmt.Printf("可以注册：%v", res)
	} else {
		fmt.Printf("不可以注册：%v", err.Content)
	}

	fmt.Printf("result: %v", res)
}