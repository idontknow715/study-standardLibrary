package main

import (
	"fmt"
	"regexp"
)

func main() {
	//s := "1435597771@qq.com"
	//fmt.Println(EmailValid(s))
	//s := "Order12345Number67890"
	//fmt.Println(GetNumber(s))
	s := "你是大傻瓜"
	fmt.Println(Replace(s))
}

func EmailValid(s string) bool {
	re := regexp.MustCompile(`\w+@\w+\.\w+`)
	return re.MatchString(s)
}

func GetNumber(s string) []string {
	re := regexp.MustCompile(`[0-9]+`)
	return re.FindAllString(s, -1)
}

func Replace(s string) string {
	re := regexp.MustCompile(`傻瓜`)
	return re.ReplaceAllString(s, "**")
}
