package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	u := User{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	data, _ := json.Marshal(&u)
	fmt.Println(string(data))

	jsonStr := `{"name": "Bob","age":30,"email":"bob@example.com"}`
	var u2 User
	json.Unmarshal([]byte(jsonStr), &u2)
	fmt.Printf("%+v\n", u2)

	// marshalIndent
	dataIndent, _ := json.MarshalIndent(u, "", " ")
	fmt.Println(string(dataIndent))
}
