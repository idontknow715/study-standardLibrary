package main

import (
	"bytes"
	"fmt"
)

// 查找类
func main() {
	fmt.Println(bytes.Contains([]byte("hello"), []byte("he")))
	fmt.Println(bytes.HasPrefix([]byte("hello"), []byte("h")))
	fmt.Println(bytes.HasSuffix([]byte("hello"), []byte("o")))
	fmt.Println(bytes.Index([]byte("hello"), []byte("l")))
	fmt.Println(bytes.LastIndex([]byte("hello"), []byte("l")))
}
