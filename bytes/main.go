package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("line\nline2\nline3\n")
	fmt.Println(string(data))

	lines := bytes.Split(data, []byte("\n"))

	var buf bytes.Buffer
	for _, line := range lines {
		buf.Write(bytes.ToUpper(line))
		buf.WriteByte('\n')
	}
	fmt.Println(string(buf.String()))
}
