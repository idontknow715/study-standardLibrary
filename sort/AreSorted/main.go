package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1, 2, 3}
	fmt.Println(sort.IntsAreSorted(num))
}
