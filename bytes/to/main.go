package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(bytes.ToUpper([]byte("go"))))
	fmt.Println(string(bytes.ToLower([]byte("LANG"))))

}
