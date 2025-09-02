package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[0-9]+`)
	rext := "hello 123 word 456"

	result := re.ReplaceAllString(rext, "#")
	fmt.Println("result", result)

	resultFunc := re.ReplaceAllStringFunc(rext, func(s string) string {
		return "[" + s + "]"
	})
	fmt.Println(resultFunc)
}
