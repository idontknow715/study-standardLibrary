package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//IoReader()
	//IoWriter()
	IoSeek()
	//io.Closer ,只要实现了Close()就是实现了io.Closer
}

// io.Reader表示可读对象,只要实现了Read(p []byte)(n int,err erros)就是Reader
func IoReader() {
	r := strings.NewReader("Hello, Go!") // string实现了io.Reader

	buf := make([]byte, 5)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Println("读到：", string(buf[:n]))
		}
		if err == io.EOF { // 读到末尾
			break
		}
	}
}

// io.Writer 表示可写对象，只要实现了Write(p []byte)(n int, err error)就是Writer
func IoWriter() {
	f, _ := os.Create("writer.txt")
	defer f.Close()

	io.WriteString(f, "hello,io.writer\n")
	fmt.Println("写入完成")
}

func IoSeek() {
	f, _ := os.Create("seek.txt")
	defer f.Close()

	f.WriteString("HelloWorld")
	// 从文件开头移动+5的位置
	f.Seek(5, 0)

	f.WriteString("GO")
	content, _ := os.ReadFile("seek.txt")
	fmt.Println("文件内容:", string(content))

}
