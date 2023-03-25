package main

import "fmt"

func main() {
	value := 3
	result := 3
	isEqual := (value == 2)

	if value == 2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	//output false



	fmt.Printf("nilai %d (%t) \n", value, isEqual) //output 3 (false)


	if value == 3 && result == 3 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	//output true


	 point := 6

	switch point  {
	case 10, 9:
		fmt.Println("perfect");
	case 8, 7:
		fmt.Println("awesome");
	default :
		fmt.Println("Not Bad")
	}

	 

}