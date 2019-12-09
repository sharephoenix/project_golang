package main

import (
	"fmt"
	"sync"
	"time"
)

//Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
//在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
//使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
//在 Lock() 之前使用 Unlock() 会导致 panic 异常
//已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
//在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
//适用于读写不确定，并且只有一个读或者写的场景

func main() {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("触发了宕机", info)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)

	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			defer func() {
				if info := recover(); info != nil {
					fmt.Println("触发了宕机**", info)
				} else {
					fmt.Println("程序正常退出**")
				}
			}()
			//panic("fatal error") /// 触发宕机
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}
}
