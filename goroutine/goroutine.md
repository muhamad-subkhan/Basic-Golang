<b>Goroutine</b>  adalah salah satu fitur utama dari bahasa pemrograman Go yang memungkinkan pembuatan kode yang berjalan secara konkuren atau paralel. Goroutine adalah unit eksekusi ringan yang dikelola oleh runtime Go. Dalam bahasa yang lebih sederhana, goroutine dapat dianggap sebagai fungsi atau prosedur yang berjalan secara independen dari program utama atau program lainnya dan berpotensi untuk berjalan secara bersamaan dengan goroutine lainnya.

Goroutine sangat berguna untuk melakukan tugas-tugas yang membutuhkan waktu yang lama seperti operasi I/O atau pemrosesan data. Dengan goroutine, kita dapat menjalankan beberapa tugas secara bersamaan, tanpa harus menunggu tugas yang satu selesai terlebih dahulu sebelum menjalankan tugas yang lain. Selain itu, goroutine juga dapat membantu dalam memanfaatkan multi-core pada mesin kita, sehingga program yang dibuat dapat berjalan lebih efisien.

Contoh

```go
package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "hallo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
```

Pada bagian main, terdapat dua pemanggilan fungsi print() dengan argumen yang berbeda. Namun, pada pemanggilan pertama, terdapat kata kunci "go" sebelum nama fungsi, yang menunjukkan bahwa fungsi tersebut akan dijalankan sebagai goroutine.

Dalam program ini, terdapat penggunaan runtime.GOMAXPROCS(2) untuk menentukan jumlah thread yang digunakan oleh Go. Pada contoh ini, terdapat dua thread yang digunakan untuk menjalankan goroutine dan fungsi print() secara konkuren.

Setelah menjalankan kedua fungsi print(), program akan menunggu input dari pengguna sebelum akhirnya berakhir. Dengan adanya penggunaan goroutine pada pemanggilan pertama fungsi print(), fungsi tersebut akan dijalankan secara konkuren dengan fungsi print() pada pemanggilan kedua, sehingga keduanya akan mencetak pesan secara bergantian.

Dalam hal ini, goroutine sangat berguna untuk menjalankan tugas-tugas yang membutuhkan waktu yang lama secara konkuren dengan tugas-tugas yang lain. Dengan demikian, program dapat berjalan lebih efisien dan dapat memanfaatkan kemampuan multi-core pada mesin kita. Namun, perlu diperhatikan juga bahwa penggunaan goroutine harus dilakukan dengan hati-hati agar tidak terjadi race condition atau deadlock pada program.
<br><br><br><br><br>



<u><b>Kenapa <i>apa kabar</i> dijalankan terlebih dahulu?</u><br>
Karena terdapat dua pemanggilan fungsi print() dengan argumen yang berbeda, yaitu "hallo" dan "apa kabar". Namun, pada pemanggilan pertama fungsi print() dengan argumen "hallo", terdapat kata kunci "go" sebelum nama fungsi, yang menunjukkan bahwa fungsi tersebut akan dijalankan sebagai goroutine.

Ketika program dijalankan, terlebih dahulu fungsi print() pada pemanggilan kedua, yaitu print(5, "apa kabar"), akan dieksekusi secara normal tanpa menggunakan goroutine. Setelah itu, baru kemudian fungsi print() pada pemanggilan pertama, yaitu go print(5, "hallo"), akan dieksekusi sebagai goroutine.

Karena goroutine memiliki karakteristik konkuren atau paralel, maka fungsi print() pada pemanggilan kedua dan goroutine pada pemanggilan pertama akan dieksekusi secara bersamaan. Sehingga, output dari program tersebut mungkin tidak selalu sama setiap kali dijalankan, dan pesan "apa kabar" pada pemanggilan kedua mungkin saja dijalankan terlebih dahulu sebelum pesan "hallo" pada goroutine dijalankan.

Hal ini terjadi karena scheduler pada Go akan memilih mana goroutine yang akan dijalankan pada saat tertentu, dan keputusan tersebut tergantung pada beberapa faktor, seperti load pada mesin dan seberapa sering goroutine tersebut melakukan blocking I/O atau pemrosesan data. Oleh karena itu, dalam penggunaan goroutine, kita harus memperhatikan penggunaannya agar tidak mengakibatkan race condition atau deadlock pada program.




<b>Referensi</b>
1. https://www.youtube.com/watch?v=oHIbeTmmTaA&feature=youtu.be&ab_channel=GolangDojo
2. https://www.youtube.com/watch?v=V-0ifUKCkBI&ab_channel=HiteshChoudhary
3. https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/