package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[;,\s]+`)
	s := "a,b;c d"
	parts := re.Split(s, -1)
	fmt.Println(parts)
}
