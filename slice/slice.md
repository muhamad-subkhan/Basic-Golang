<b>Slice</b> <br>
tipe data yang digunakan untuk merepresentasikan sekumpulan elemen yang terorganisir secara linier. Slice mirip dengan array, namun memiliki fleksibilitas yang lebih besar karena ukuran slice dapat berubah secara dinamis dan tidak perlu didefinisikan pada saat kompilasi.

contoh 
```go
package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

<b>Primes</b> -> nama variable yang menampung array

Kemudian, terdapat variabel s yang merupakan slice yang diambil dari array primes dimulai dari indeks ke-1 (yaitu nilai 3) hingga indeks ke-3 (yaitu nilai 7), dengan menggunakan operator :.

Setelah itu, dilakukan pencetakan nilai s menggunakan fmt.Println(s), sehingga hasil yang dikeluarkan adalah [3 5 7] yang merupakan slice dari nilai 3 hingga 7 dari array primes.


<b>Referensi</b>
1. https://www.w3schools.com/go/go_slices.php
2. https://www.youtube.com/watch?t=6473&v=YS4e4q9oBaU&feature=youtu.be&ab_channel=freeCodeCamp.org