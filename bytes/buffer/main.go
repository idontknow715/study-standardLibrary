package main

import (
	"bytes"
	"fmt"
)

// Buffer是一个可变的字节缓冲区，最常用在字符串拼接、数据收集、IO读写场景
func main() {
	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.Write([]byte(" world"))
	fmt.Println(buf.String())

	// 当作reader
	bufr := bytes.NewBufferString("hello world")
	data := make([]byte, 5)
	bufr.Read(data)
	fmt.Println(string(data))
}
