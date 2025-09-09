package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("任务", id, "完成")
		}(i)
	}
	wg.Wait()
	fmt.Println("所有任务完成")
}
