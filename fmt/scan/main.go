package main

import "fmt"

func main() {
	//Scan()
	Scanln()
}

// 用来读取多个输入项，可以是不同类型
func Scan() {
	var age int
	fmt.Scan(&age)
	fmt.Printf("age: %d\n", age)
}

// 用来从标准输入读取一行数据
func Scanln() {
	var name string
	fmt.Scanln(&name)
	fmt.Printf("name: %s\n", name)
}
