package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)

	match := re.FindStringSubmatch("myemail@test.com")
	fmt.Println("match: ", match)

	matches := re.FindAllStringSubmatch("a@b.com c@d.org", -1)
	for _, m := range matches {
		fmt.Println("m: ", m)
	}
}
