package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取环境变量
	path := os.Getenv("PATH")
	fmt.Println("PATH =", path)

	// 设置环境变量
	os.Setenv("MY_VAR", "helloworle")
	fmt.Println("MY_VAR", os.Getenv("MY_VAR"))

	// 获取所有环境变量
	envs := os.Environ()
	for _, e := range envs[:5] { // 只打印前5个
		fmt.Println(e)
	}
}
