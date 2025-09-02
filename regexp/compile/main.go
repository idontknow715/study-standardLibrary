package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Println("正则错误：", err)
	}
	fmt.Println(re1.MatchString("123abc"))

	re2 := regexp.MustCompile(`\w+`)
	fmt.Println(re2.MatchString("hello"))
}
