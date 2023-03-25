package main

import "fmt"

func main() {

	var a int = 42
	var p *int = &a

	fmt.Println("Nilai a: ", a)
	fmt.Println("Alamat a: ", &a)
	fmt.Println("Nilai p: ", p)
	fmt.Println("Nilai yang ditunjuk oleh p: ", *p)

	*p = 21
	fmt.Println("Nilai a setelah diubah melalui p: ", a)

}
