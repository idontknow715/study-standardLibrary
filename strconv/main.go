package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// go run ./main.go 2
	fmt.Println("strconv标准库学习小任务,命令行参数解析")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}
	numstr := os.Args[1]
	num, err := strconv.Atoi(numstr)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}
	fmt.Println("你输入的数字平方是: %d\n", num*num)
}
