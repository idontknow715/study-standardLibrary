package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05", "2025-08-20 23:02:00")

	if err != nil {
		fmt.Println("解析时间错误:", err)
	}
	fmt.Println("解析的时间为：", t)
}
