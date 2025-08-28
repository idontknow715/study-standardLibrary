package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//IoCopy()
	//IoReadAll()
	//IoReadFull()
	IoWriteString()
}

// io.Copy 把数据从一个Reader拷贝到Writer中
func IoCopy() {
	src, _ := os.Open("seek.txt")
	defer src.Close()

	dst, _ := os.Create("copy.txt")
	defer dst.Close()

	n, _ := io.Copy(dst, src)
	fmt.Println("拷贝字节数：", n)
}

// io.ReadAll 读取一个Reader的所有内容
func IoReadAll() {
	r := strings.NewReader("Hello ReadAll")
	data, _ := io.ReadAll(r)
	fmt.Println("结果为:", string(data))
}

// io.ReadFull 保证读取制定长度，否则报错
func IoReadFull() {
	r := strings.NewReader("abcdefg")
	buf := make([]byte, 4)

	_, err := io.ReadFull(r, buf)
	if err != nil {
		fmt.Println("读取失败：", err)
		return
	}
	fmt.Println("固定长度结果：", string(buf))
}

// 快速把字符串写入Writer
func IoWriteString() {
	f, err := os.OpenFile("writer.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer f.Close()

	n1, err := io.WriteString(f, "第一行日志\n")
	if err != nil {
		fmt.Println("写入失败,err: ", err)
	} else {
		fmt.Println("写入的长度为：", n1)
	}
	io.WriteString(f, "第二行日志\n")
}
