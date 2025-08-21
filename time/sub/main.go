package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	later := now.Add(2 * time.Hour)
	fmt.Println("两小时后：", later)

	// 计算时间差
	diff := later.Sub(now)
	fmt.Println("相差：", diff.Hours(), "小时")
}
