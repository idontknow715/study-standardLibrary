package main

import (
	"fmt"
	"sort"
)

func main() {
	sortInt := []int{3, 2, 4}
	sort.Ints(sortInt)
	fmt.Println("sortInt: ", sortInt)

	sortFloat := []float64{3.1, 3.2, 3.0}
	sort.Float64s(sortFloat)
	fmt.Println("sortFloat: ", sortFloat)

	// 按照字典序列排序,第一个不同的字符决定排序顺序G：71 P：80 C：67
	sortString := []string{"Go", "C", "Python"}
	sort.Strings(sortString)
	fmt.Println("sortString: ", sortString)
}
