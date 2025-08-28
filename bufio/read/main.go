package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("bufio/test.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("读取失败：", err)
			break
		} else if err == io.EOF {
			fmt.Println("读完了")
			break
		}
		fmt.Println("读取到的内容为： ", line)
	}
}
