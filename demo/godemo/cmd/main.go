package main

import (
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
	Flag bool `json:"flag,omitempty"`
}

//golang chan<- 和 <-chan，作为函数参数时
//开始时看到这个实在没明白怎么回事
//
//测试了下才知道原来
//
//<-chan int  像这样的只能接收值
//
//chan<- int  像这样的只能发送值

func sendchan(s chan<- int)  {
	s <- 100
}

func receive(s <-chan int) {
	a := <-s
	fmt.Println( "a",a)
}

func main() {

<<<<<<< HEAD:demo/godemo/cmd/main.go
	var user User
	user.Name = "alexluan"
	user.Address = "shanghai"
	user.Flag = false
	jbyt, _ := json.Marshal(user)
=======
	myNum := []int{10, 20, 30, 40, 50}
	// 创建新的切片，其长度为 2 个元素，容量为 4 个元素
	newNum := myNum[1:3]
	//myNum[1] = 99999
	copy(myNum, newNum)
	//newNum = append(newNum, 88)
	//newNum = append(newNum, 80)
	//newNum = append(newNum, 81)
	//newNum = append(newNum, 82)
	//newNum[2] = 9
	fmt.Println(newNum, myNum)

	//async := sync.WaitGroup{}
	//chans := make(chan int, 5)
	//chans <- 100
	//chans <- 3
	//chans <- 5
	//chans <- 6
	//async.Add(1)
	//go func() {
	//	for {
	//		v := <- chans
	//		switch v {
	//		case 3:
	//			fmt.Println("3333333")
	//			break
	//		case 100:
	//			fmt.Println("100-100")
	//			break
	//		case 5:
	//			fmt.Println("5555-555")
	//			break
	//		default:
	//			fmt.Println("nnnnnullll")
	//			async.Done()
	//			break
	//		}
	//	}
	//}()
	//async.Wait()

	//var newMap map[string]string
	//defer func() {
	//
	//
	//	if ok := recover(); ok != nil {
	//		fmt.Println("0000000k", ok)
	//	}
	//}()
	//newMap = make(map[string]string)
	//newMap["a"] = "aaaaa"
	//newMap["b"] = "bbbbbb"
	////delete(newMap, "a")
	//for key, v := range newMap {
	//	fmt.Println(key, v)
	//}


	//arr := []int{1,2,3,4,5,6,7}
	//for key := range arr {
	//	fmt.Println(key)
	//}
	//arr = append(arr, []int{8,9}...)
	//fmt.Println(arr[0:])

	//var user User
	//user.Name = "alexluan"
	//user.Address = "shanghai"
	//jbyt, _ := json.Marshal(user)
	//fmt.Println(string(jbyt))


>>>>>>> f9b284fa7253a6e4ba7f7cc06b73c8ecea2aa35b:godemo/cmd/main.go

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