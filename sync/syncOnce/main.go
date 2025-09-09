package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("只会执行一次的初始化逻辑")
		})
		fmt.Println("goroutine ", i, "执行了Do()")
	}
}
