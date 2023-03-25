package main

import "fmt"

func main() {

	type identity struct {
		name string
		age  int
	}
	identitas  := identity{name: "John doe", age: 25}

	// fmt.Println(identitas)
	fmt.Println(identitas.name)
	fmt.Println(identitas.age)
}