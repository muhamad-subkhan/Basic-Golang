<b>Function</b><br>
blok kode yang dapat dipanggil oleh bagian lain dari program untuk melakukan tugas tertentu. Fungsi memungkinkan Anda untuk mengorganisir kode program Anda menjadi bagian yang lebih kecil dan lebih mudah diatur, serta memungkinkan Anda untuk mengeksekusi kode yang sama beberapa kali tanpa harus mengetiknya ulang setiap kali.

<b>Contoh</b>

```go
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
```

Kode diatas adalah contoh penggunaan function dengan parameter dan memiliki nilai kembalian/return string

<b>func name(first string) string</b> -> name adalah nama function yang berguna untuk menyimpan baris code dan untuk memanggil/mengeksekusi code yang berada di dalam function name.

<b>first string</b> -> adalah parameter yang digunakan sebagai inputan. dalam contoh diatas, first memiliki tipe data string. yang artinya first hanya bisa di isi dengan data yang bertipe string.



<b>Referensi</b><br>
1. https://www.golangprograms.com/go-language/functions.html