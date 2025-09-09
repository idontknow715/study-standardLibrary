package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//mutex()
	//rwMutex()
	//waitGroup()
	//once()
	condfunc()
}

// 并发安全计数器
func mutex() {
	var mu sync.Mutex
	val := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			val++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("val: ", val)
}

// 并发安全词典
func rwMutex() {
	data := make(map[string]string)
	var rw sync.RWMutex
	var wg sync.WaitGroup

	// 写
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rw.Lock()
			key := fmt.Sprintf("k%d", i)
			val := fmt.Sprintf("v%d", i)
			data[key] = val
			fmt.Println("写入:", key, val)
			rw.Unlock()
		}(i)
	}
	wg.Wait()

	// 读
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("k%d", i)
			rw.RLock()
			val := data[key]
			fmt.Println("读取：", key, val)
			rw.RUnlock()
		}(i)
	}
	wg.Wait()
}

func waitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("my id is:", id)
		}(i)

	}
	wg.Wait()
	fmt.Println("done")
}

func once() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("第%d次\n", id)
			once.Do(func() {
				fmt.Println(id)
			})
		}(i)

	}
	wg.Wait()
}

func condfunc() {
	cond := sync.NewCond(&sync.Mutex{})
	// 模拟信号
	ready := false

	// 消费者
	go func() {
		cond.L.Lock()
		for !ready {
			fmt.Println("消费者等待中...")
			cond.Wait()
		}
		fmt.Println("消费")

		cond.L.Unlock()
	}()

	// 模拟生产者
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	// 生产数据
	ready = true
	cond.Signal()
	cond.L.Unlock()
	time.Sleep(1 * time.Second)
}
