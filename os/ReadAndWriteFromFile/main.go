package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("data.txt")

	// 写入数据
	f.Write([]byte("hello go \n"))
	f.WriteString("写入字符串\n")

	// 移动文件指针到开头
	//whence = 0 从开头 1 从当前位置 2 从末尾
	f.Seek(0, 0)

	// 读取内容
	buf := make([]byte, 20)
	n, _ := f.Read(buf)
	fmt.Println("读取到的内容：", string(buf[:n]))

	// 关闭文件
	f.Close()
}
