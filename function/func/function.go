package main

import "fmt"

func main() {
	var first string = "John"

	fmt.Println(name(first))

}

func name(first string) string {
	var result string

	var last string
	last = "doe!"

	result = first + " " + last

	return result
}
