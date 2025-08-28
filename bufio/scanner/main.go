package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入内容（输入exit退出）:")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}
		fmt.Println("你输入了：", text)
	}
}
