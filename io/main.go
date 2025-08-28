package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//IoCopy()
	Std()
}

func IoCopy() {
	src, _ := os.Open("writer.txt")
	defer src.Close()

	dst, _ := os.Create("output.txt")
	defer dst.Close()

	n, _ := io.Copy(dst, src)
	fmt.Println("拷贝的长度:", n)
}

func Std() {
	input, err := os.Create("io/input.txt")
	if err != nil {
		fmt.Println("创建文件失败,err: ", err)
	}
	fmt.Println("请输入一行内容: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()

		// 写入文件
		_, err := input.WriteString(text + "\n")
		if err != nil {
			fmt.Println("写入失败：", err)
			return
		}
		fmt.Println("写入成功！")
	}
}
