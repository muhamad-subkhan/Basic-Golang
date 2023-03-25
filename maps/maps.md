<b>Maps</b> <br>
tipe data yang digunakan untuk merepresentasikan kumpulan data dengan pasangan key-value yang tidak memiliki urutan tertentu. Mirip dengan array dan slice, maps juga digunakan untuk menyimpan data, namun maps memiliki fleksibilitas yang lebih tinggi karena key-nya dapat berupa tipe data apa saja, termasuk tipe data yang telah didefinisikan oleh pengguna.

contoh 


```go
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
```

Pada program di atas, terdapat variabel maps myMap yang diinisialisasi dengan tipe data key string dan tipe data value int. Kemudian, terdapat beberapa operasi yang dilakukan pada maps, antara lain:

1. Menambahkan key-value pairs ke maps dengan operator [ ] pada baris ke-8 hingga ke-10.
2. Mengakses nilai pada key tertentu dengan menggunakan operator [ ] pada baris ke-13 dan ke-14.
3. Mengubah nilai pada key tertentu dengan menggunakan operator [ ] pada baris ke-16.
4. Menghapus key-value pairs dari maps dengan fungsi delete pada baris ke-19.
5. Mengecek apakah suatu key ada dalam maps dengan menggunakan operasi if exists pada baris ke-22 hingga ke-26.

<b>referensi</b> <br>
1. https://www.w3schools.com/go/go_maps.php
2. https://www.youtube.com/watch?v=yJE2RC37BF4&ab_channel=TechWithTim