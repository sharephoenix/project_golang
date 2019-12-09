package maproutine

import (
	"fmt"
	"strconv"
	"sync"
)

type M struct {
	Map map[string]string
}

// Set ...
func (m *M)Set(k string, v string) {
	m.Map[k] = v
}

// Get ...
func (m *M) Get(key string) string {
	return m.Map[key]
}

// TestMap  ...
func TestMap() {
	c := M{Map: make(map[string]string)}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
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