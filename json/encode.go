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

	object := []User{{"John Doe", 25}, {"Jhonny", 25}}
	jsonData, err := json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonString := string(jsonData)
	fmt.Println(jsonString)


}