package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05", "2025-12-31 23:59:59")
	if err != nil {
		fmt.Println("解析时间错误")
	}
	fmt.Println(t)
}
