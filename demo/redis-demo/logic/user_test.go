package logic

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	account        = "18817322819"
	password       = "qwe123"
	name           = "alexlaun"
	address        = "shanghai"
	sex      int64 = 0
)

func Test_Register_test(t *testing.T) {
	res, canRegister := CheckRegister(account)
	if !canRegister {
		fmt.Printf("用户存在，不可以注册：%v\n", res)
		return
	} else {
		fmt.Printf("可以注册：%v\n", res)
	}
	fmt.Println(account, password, name, address, sex)
	err0 := Register(account, password, name, address, sex)
	if err0 == nil {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败", err0)
	}
}

func Test_CheckRegister(t *testing.T) {
	res, canRegister := CheckRegister(account)
	if !canRegister {
		fmt.Printf("不可以注册：%v\n", res)
	} else {
		fmt.Printf("可以注册：%v\n", res)
	}

	fmt.Printf("result: %v\n", res)
}

func Test_Login(t *testing.T) {
	user := Login(account, password)
	if user == nil {
		fmt.Println("登录失败")
	} else {
		jsonBytes, _ := json.Marshal(&user)
		fmt.Printf("登录成功:%v", string(jsonBytes))
	}
}
