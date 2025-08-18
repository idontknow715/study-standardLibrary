package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(bytes.Replace([]byte("foo foo"), []byte("foo"), []byte("bar"), 2)))
}
