package main

import (
	"bytes"
	"fmt"
)

func main() {
	parts := bytes.Split([]byte("a,b,c"), []byte(","))
	fmt.Println("part: ", parts)

	joins := bytes.Join(parts, []byte("-"))
	fmt.Println("joins: ", string(joins))
}
