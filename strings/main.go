package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, Golang!"

	// 是否包含
	fmt.Println(strings.Contains(s, "Go"))

	// 前缀判断
	fmt.Println(strings.HasPrefix(s, "Hello"))

	// 后缀判断
	fmt.Println(strings.HasSuffix(s, "!"))

	// 子串位置
	fmt.Println(strings.Index(s, "Go"))

	// 替换 n：表示替换几次，比如n为2，字符串有两个Go，那么都会被替换，如果n为1，那么就只会替换一个Go
	fmt.Println(strings.Replace(s, "Go", "Java", 1))

	// 拆分
	fmt.Println(strings.Split(s, ","))

	// 拼接
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	// 转小写
	fmt.Println(strings.ToLower(s))

	// 转大写
	fmt.Println(strings.ToUpper(s))
	s1 := []string{"go", "lang", "rocks"}
	fmt.Println(test(s1))
}

func test(s []string) string {
	s1 := strings.Join(s, "-")
	s2 := strings.ToUpper(s1)
	return s2
}
