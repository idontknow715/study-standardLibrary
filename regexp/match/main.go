package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[0-9]+`)

	// matchString
	fmt.Println(re.MatchString("abc123c"))
	fmt.Println(re.MatchString("abc"))

	// FindString
	fmt.Println(re.FindString("abc123xyz"))

	// FindAllString
	fmt.Println(re.FindAllString("a1b22c333", -1))
}
