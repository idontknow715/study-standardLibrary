package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("当前进程id：", os.Getpid())
	fmt.Println("父进程id：", os.Getppid())
}
