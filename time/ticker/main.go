package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 延迟执行
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("2秒后")
	})

	// 2. 周期执行：ticker
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("每秒执行一次")
	}
}
