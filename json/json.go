package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {

	var jsonString = `{"Name": "John Doe", "Age": 25}`
	var jsonData = []byte(jsonString)
	var data User

	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Name :", data.FullName)
	fmt.Println("Age :", data.Age)
}
