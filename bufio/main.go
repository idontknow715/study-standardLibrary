package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	//CopyToOutput()
	//Log("你好")
	//fmt.Println(ScannerWord())
	FilterError()
}

func CopyToOutput() {
	file, err := os.Open("bufio/input.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	defer file.Close()
	output, err := os.Create("bufio/output.txt")
	if err != nil {
		fmt.Println("创建文件失败:", err)
	}
	defer output.Close()
	writer := bufio.NewWriter(output)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("读取失败：", err)
			break
		} else if err == io.EOF {
			fmt.Println("读完了")
			if len(line) > 0 {
				writer.WriteString(line)
			}
			break
		}
		_, err = writer.WriteString(line)
		if err != nil {
			fmt.Println("写入到output.txt失败:", err)
		}

	}
	writer.Flush()
}

func Log(message string) {
	file, err := os.OpenFile("bufio/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err = writer.WriteString(fmt.Sprintf("[%s] %s \n", now, message))
	if err != nil {
		fmt.Println("写入失败：", err)
	}
}

func ScannerWord() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入单词（输入exit退出）:")
	num := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		} else {
			num++
		}

	}
	return num
}

func FilterError() {
	in, err := os.Open("bufio/log.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer in.Close()

	out, err := os.OpenFile("bufio/error.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开error文件失败:", err)
		return
	}
	defer out.Close()

	reader := bufio.NewReader(in)

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if len(line) > 0 && strings.Contains(line, "ERROR") {
			_, writeErr := writer.WriteString(line)
			if writeErr != nil {
				fmt.Println("写入失败：", writeErr)
				break
			}
		}

	}
}
