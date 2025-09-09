package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Println("goroutine ", id, "把counter 改为", counter)
			mu.Unlock()

		}(i)
	}
	wg.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Println("最终counter为：", counter)
}
