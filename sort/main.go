package main

import (
	"fmt"
	"sort"
)

func main() {
	numInt := []int{5, 2, 9, 1, 5, 6}
	sortInt(numInt)

	fmt.Println(numInt)

	aresort := []int{1, 2, 4, 3}
	fmt.Println(areSort(aresort))

	numString := []string{"banana", "apple", "pear"}
	fmt.Println(sortString(numString))

	numMySort := []Person{
		{"A", 19},
		{"B", 8},
	}
	fmt.Println(mySort(numMySort))
}

// 给整数切片排序
func sortInt(num []int) []int {
	sort.Ints(num)
	return num
}

// 把字符串切片排序

func sortString(num []string) []string {
	sort.Strings(num)
	return num
}

// 判断是否有序
func areSort(num []int) bool {
	return sort.IntsAreSorted(num)
}

type Person struct {
	Name string
	Age  int
}

func mySort(num []Person) []Person {
	sort.Slice(num, func(i, j int) bool {
		return num[i].Age < num[j].Age
	})
	return num
}
