package main

import (
	"encoding/json"
	"os"
)

func main() {
	file, _ := os.Create("encodingJson/user.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	encoder.Encode(map[string]string{"hello": "world"})
}
