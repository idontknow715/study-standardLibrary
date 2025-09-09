package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	//Tag()
	//NewDecoder()
	NewEncoder()
}

func MarshalAndUnMarshal() {
	u1 := User{
		Name:  "bob",
		Age:   1,
		Email: "bob@example.com",
	}
	data, _ := json.Marshal(&u1)
	fmt.Println(string(data))

	var uDecoder User
	json.Unmarshal(data, &uDecoder)
	fmt.Println("uDecoder: ", uDecoder)
}

func Tag() {
	type userTag struct {
		name  string `json:"name"`
		age   int    `json:"age,omitempty"`
		email string `json:"-"`
	}
	usertag := userTag{name: "Alice"}
	data, _ := json.Marshal(usertag)
	fmt.Println(string(data))
}

func NewDecoder() {
	file, err := os.Open("encodingJson/config.json")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()
	type config struct {
		Port  int    `json:"port"`
		Db    string `json:"db"`
		Debug bool   `json:"debug"`
	}
	de := json.NewDecoder(file)
	var configdata config
	err = de.Decode(&configdata)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(configdata)
}

func NewEncoder() {
	file, _ := os.OpenFile("encodingJson/log.json", os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0644)
	defer file.Close()

	type LogEntry struct {
		Time string `json:"time"`
		Msg  string `json:"msg"`
	}
	log := LogEntry{
		Time: time.Now().String(),
		Msg:  "test",
	}
	en := json.NewEncoder(file)
	en.SetIndent("", " ")
	err := en.Encode(&log)
	if err != nil {
		fmt.Println("err: ", err)
	}
}
