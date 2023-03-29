package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)
	sayHello := func (who string) {
		var data = fmt.Sprintf("Hallo %s", who)
		message <- data;
	}


	go sayHello("John Doe")
	go sayHello("Jhonny")
	go sayHello("kitty")
	
	messageSatu := <- message
	fmt.Println(messageSatu)

	messageDua := <- message
	fmt.Println(messageDua)

	messageTiga := <- message
	fmt.Println(messageTiga)

}