package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("halo")
	fmt.Println("selamat datang")
	os.Exit(1) // matikan os.exit dengan cara cammad //os.exit untuk melihat perbedaannya
	fmt.Println("lagi")
}