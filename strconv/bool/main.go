package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatBool(true))

	//ParseBool 支持："1","t","T","true","TRUE","True"都表示true
	b, _ := strconv.ParseBool("t")
	fmt.Println("b: ", b)
}
