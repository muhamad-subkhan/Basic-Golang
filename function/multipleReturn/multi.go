package main


import "fmt"


func main() {
	add, sub := calculate(10, 5)
    fmt.Println("Addition result:", add)
    fmt.Println("Subtraction result:", sub)
}


func calculate(a int, b int) (int, int) {
	return a + b, a - b
}
