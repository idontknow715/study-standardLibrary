package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("encodingJson/user.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data map[string]string
	decoder.Decode(&data)
	fmt.Println("data: ", data)
}
