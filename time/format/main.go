package main

import (
	"fmt"
	"time"
)

// Go的时间格式化是用固定的参数时间：2006-01-02 15:04:05
func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05"))
}
