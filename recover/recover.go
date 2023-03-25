package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error Occured", r)
	} else {
		fmt.Println("Aplication Successfully Running")
	}
}


func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == ""{
		return false, errors.New("Kata Tidak Boleh Kosong")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		defer fmt.Println("tetap terbaca")
		panic(err.Error())
	}
}