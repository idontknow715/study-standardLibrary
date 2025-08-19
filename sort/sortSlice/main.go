package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int64
	}{
		{"Alice", 19},
		{"BOB", 17},
		{"Candy", 22},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	fmt.Println("people is: ", people)
}
