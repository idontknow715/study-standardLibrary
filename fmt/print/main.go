package main

import "fmt"

func main() {
	fmt.Println("Hello,World") // 输出并换行
	fmt.Print("Go is awesome") // 输出不换行

	name := "Go"
	age := 10
	pi := 3.14159
	// 使用格式化符号
	fmt.Printf("Language: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Pi: %.2f\n", pi)
}
