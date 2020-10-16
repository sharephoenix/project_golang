package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // etcd 集群三个实例的端口
		DialTimeout: 2 * time.Second,
	})
	wg.Add(1)
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")

	defer cli.Close()

	go func() {
		for true {
			rch := cli.Watch(context.Background(), "name", clientv3.WithPrevKV()) //阻塞在这里，如果没有key里没有变化，就一直停留在这里
			for wresp := range rch {
				for _, ev := range wresp.Events {
					fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				}
			}
			fmt.Println("for")
		}
	}()

	go func() {
		for true {
			rch := cli.Watch(context.Background(), "address", clientv3.WithPrevKV()) //阻塞在这里，如果没有key里没有变化，就一直停留在这里
			for wresp := range rch {
				for _, ev := range wresp.Events {
					fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				}
			}
			fmt.Println("for")
		}
	}()

	wg.Wait()
}
