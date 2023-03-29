package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonString := `{"Name": "John Doe", "Age": 25}`
	jsonData := []byte(jsonString)

	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)
	fmt.Println("User:", data["Name"])
	fmt.Println("Age:", data["Age"])

}