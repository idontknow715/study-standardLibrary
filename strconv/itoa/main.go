package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1. 数字转字符串
	// int转字符串
	fmt.Println(strconv.Itoa(123))
	//int64 转指定进制的字符串
	fmt.Println(strconv.FormatInt(123, 16))
	//uint转字符串
	fmt.Println(strconv.FormatUint(123, 16))

	//浮点数转字符串
	//strconv.FormatFloat(f float64,fmt byte,prec,bitSize int)string
	// fmt参数常用:'f'小数,'e'科学计数法,'g' 自动选择f还是e
	// prec 表示保留小数位数，-1表示尽可能精确
	//bitSize 一般32或者64
	fmt.Println(strconv.FormatFloat(3.14159, 'f', 4, 64))

}
