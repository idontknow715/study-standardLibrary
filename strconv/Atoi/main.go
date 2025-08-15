package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("123")
	fmt.Println("i: ", i)

	// base 表示进制（2～36）
	//bitSize 表示目标精度 32/64
	n, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println("n: ", n)

	u, _ := strconv.ParseUint("ff", 16, 64)
	fmt.Println("u: ", u)

	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("f: ", f)
}
