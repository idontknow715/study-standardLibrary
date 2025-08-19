package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int64
}

type ByName []Person

func main() {
	b := ByName{
		{"Alice", 12},
		{"Charlie", 15},
		{"Bob", 17},
	}
	sort.Sort(b)
	fmt.Println("b is: ", b)
}

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Less(i, j int) bool {
	return b[i].Name >= b[j].Name
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
