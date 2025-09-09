package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		data = make(map[string]string)
		rw   sync.RWMutex
	)
	rw.Lock()
	data["name"] = "Go"
	rw.Unlock()
	var wg sync.WaitGroup
	// 同时进行一些写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rw.Lock()
			data["name"] = fmt.Sprintf("Go%d", i)
			rw.Unlock()
		}(i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rw.RLock()
			fmt.Println("goroutine ", id, "读取到的name = ", data["name"])
		}(i)
	}

	wg.Wait()
}
