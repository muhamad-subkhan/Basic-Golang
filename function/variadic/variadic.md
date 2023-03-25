<b>Function Variadic</b><br>
kemampuan suatu fungsi untuk menerima jumlah argumen yang tidak terbatas. Fungsi variadic memungkinkan kita untuk memanggil fungsi dengan jumlah argumen yang bervariasi, dan kita dapat menangani argumen-argumen tersebut secara dinamis.


```go
package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers)) // length
	return avg
}
```

<b>Penjelasan</b><br>
Fungsi `calculate` menerima parameter variadic `numbers` dengan tipe data `int`, yang berarti fungsi ini dapat menerima sejumlah argumen yang bervariasi. Di dalam fungsi `calculate`, kita menggunakan perulangan `for` untuk mengiterasi seluruh argumen yang diberikan, dan menambahkan nilai-nilai tersebut ke dalam variabel `total`.

Setelah seluruh nilai dijumlahkan, kita menghitung rata-rata dengan membagi jumlah total dengan jumlah argumen yang diberikan, dan hasilnya disimpan ke dalam variabel `avg` dengan tipe data `float64`.

Di dalam fungsi `main`, kita memanggil fungsi `calculate` dengan memberikan sejumlah argumen bilangan. Fungsi `calculate` kemudian akan menghitung rata-rata dari semua bilangan tersebut, dan mengembalikan nilai rata-rata dalam bentuk `float64`. Hasil rata-rata tersebut kemudian dicetak ke konsol menggunakan fungsi `Printf` dari package "fmt".

Dengan menggunakan fungsi variadic pada Go, kita dapat membuat fungsi yang lebih fleksibel dan dapat menangani sejumlah argumen yang bervariasi, seperti pada contoh di atas yang dapat menghitung rata-rata dari sejumlah bilangan yang diberikan.













<b>Referensi</b><br>
1. https://www.scaler.com/topics/golang/variadic-functions-in-go/