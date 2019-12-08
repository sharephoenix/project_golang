package maproutine

import (
	"fmt"
	"strconv"
	"sync"
)

type SM struct {
	Map    map[string]string
	lock sync.RWMutex // 加锁
}

// Set ...
func (m *SM) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = value
}

// Get ...
func (m *SM) Get(key string) string {
	// 特意加锁，不然有机会出错
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.Map[key]
}

// TestSyncMap  ...
func TestSyncMap() {
	c := SM{Map: make(map[string]string)}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			c.Set(k, v)
			fmt.Println("k=:%v,v:=%v\n", k, c.Get(k))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("ok finished.")
}
