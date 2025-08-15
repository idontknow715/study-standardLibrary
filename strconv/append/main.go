package main

import (
	"fmt"
	"strconv"
)

func main() {
	//AppendXxx 方法直接往[]byte追加字符串表示，减少内存分配（性能优化时使用）
	//strconv.AppendInt(dst []byte, i int64, base int) []byte
	//strconv.AppendBool(dst []byte, b bool) []byte
	//strconv.AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
	buf := []byte("value: ")
	buf = strconv.AppendInt(buf, 123, 10)
	fmt.Println(string(buf))
}
