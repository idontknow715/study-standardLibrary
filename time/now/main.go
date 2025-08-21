package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间：", now)
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())
	fmt.Println("日:", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())
}
