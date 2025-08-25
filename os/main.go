package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//ReadAll("data.txt")
	//WriteLog("程序启动")
	//WriteLog("执行某个操作成功")
	SetEnv("xxx")
}

func ReadAll(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	// 用缓冲区
	//buf := make([]byte, 1024)
	//f.Read(buf)
	//fmt.Println("读取到的所有内容为：", string(buf))
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("读取失败")
	}
	fmt.Println("读取到的所有内容为：", string(data))
	f.Close()

}

func WriteLog(message string) {
	//打开文件，如果不存在就创建，存在就追加
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败：", err)
		return
	}
	defer f.Close()

	now := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] %s\n", now, message)
	if _, err := f.WriteString(logLine); err != nil {
		fmt.Println("写入日志失败:", err)
	}
}

func SetEnv(value string) {
	path := os.Getenv("PATH")
	fmt.Println("path: ", path)

	err := os.Setenv("MYNAME", value)
	if err != nil {
		fmt.Println("设置环境变量失败: ", err)
		return
	}

	fmt.Printf("设置的环境变量为: %s=%s", "MYNAME", value)
}
