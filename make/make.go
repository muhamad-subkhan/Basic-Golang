package main

import "fmt"

func main() {

	var fruits = make([]string, 2)
	fruits[0] = "foo"
	fruits[1] = "bar"

	fmt.Println(fruits)
}