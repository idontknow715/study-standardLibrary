package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	ready := false

	// 消费者
	go func() {
		cond.L.Lock()
		for !ready {
			fmt.Println("消费者等待中。。。")
			cond.Wait()
		}
		fmt.Println("消费者收到信号，开始消费！")
		cond.L.Unlock()
	}()
	// 模拟生产者
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	ready = true
	fmt.Println("生产者好了，通知消费者")
	cond.Signal()
	cond.L.Unlock()
	time.Sleep(1 * time.Second)
}
