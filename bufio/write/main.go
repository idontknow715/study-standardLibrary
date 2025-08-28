package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("bufio/log.txt")
	if err != nil {
		fmt.Println("创建文件失败：", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i <= 10; i++ {
		writer.WriteString(fmt.Sprintf("第%d行日志\n", i))
	}
	writer.Flush()
}
