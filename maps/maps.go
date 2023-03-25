package main

import "fmt"

func main() {

	    // Membuat variabel maps baru
		var myMap map[string]int = make(map[string]int)

		// Menambahkan key-value pairs ke maps
		myMap["satu"] = 1
		myMap["dua"] = 2
		myMap["tiga"] = 3
	
		// Mengakses nilai pada key tertentu
		fmt.Println("Nilai pada key \"satu\" adalah", myMap["satu"])
		fmt.Println("Nilai pada key \"dua\" adalah", myMap["dua"])
	
		// Mengubah nilai pada key tertentu
		myMap["satu"] = 10
		fmt.Println("Nilai pada key \"satu\" setelah diubah adalah", myMap["satu"])
	
		// Menghapus key-value pairs dari maps
		delete(myMap, "tiga")
		fmt.Println("Setelah menghapus key \"tiga\", maka variabel myMap menjadi", myMap)
	
		// Mengecek apakah suatu key ada dalam maps
		value, exists := myMap["satu"]
		if exists {
			fmt.Println("Key \"satu\" terdapat pada maps dan bernilai", value)
		} else {
			fmt.Println("Key \"satu\" tidak terdapat pada maps")
		}
}