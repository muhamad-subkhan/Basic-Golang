<b>Wait Group</b><br>
tipe data pada bahasa pemrograman Go yang digunakan untuk menunggu sekelompok goroutine selesai menjalankan tugasnya sebelum melanjutkan eksekusi program ke tahap selanjutnya. WaitGroup berfungsi sebagai penghitung sederhana yang dapat digunakan untuk menunggu satu atau beberapa goroutine menyelesaikan tugasnya sebelum melakukan tindakan selanjutnya.

Cara kerja WaitGroup cukup sederhana. Pada awalnya, WaitGroup diberikan nilai awal yang menunjukkan jumlah goroutine yang akan dijalankan. Setiap goroutine akan menambahkan jumlah penghitung WaitGroup sebelum selesai menjalankan tugasnya. Ketika semua goroutine telah menyelesaikan tugasnya, penghitung WaitGroup menjadi 0, dan program dapat melanjutkan eksekusinya ke tahap selanjutnya.

`sync.WaitGroup` biasanya digunakan dalam aplikasi Go yang bersifat concurrent, dimana beberapa goroutine perlu berjalan secara simultan, tetapi pada akhirnya harus menunggu satu sama lain sebelum melakukan tindakan selanjutnya. Penggunaan `sync.WaitGroup` akan memastikan bahwa program menunggu semua goroutine selesai sebelum melanjutkan eksekusi ke tahap selanjutnya, sehingga menghindari kondisi race condition atau deadlock.
<br><br><br>


<b>Contoh Code</b><br>


```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {

	defer wg.Done()
	fmt.Println(message)
}



func main() {

	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("data %d\n", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()
}
```
Proses berjalannya kode di atas adalah sebagai berikut:

1. Pertama-tama, program mengatur jumlah maksimum prosesor yang dapat digunakan dengan menggunakan `runtime.GOMAXPROCS(8)`. Dalam hal ini, program akan menggunakan maksimum 8 prosesor.

2. Selanjutnya, program membuat sebuah `sync.WaitGroup` yang disimpan dalam variabel `wg`.

3. Program melakukan pengulangan sebanyak 5 kali menggunakan for loop dengan variabel `i` sebagai penghitung. Dalam setiap perulangan, program membuat sebuah data dengan menggunakan `fmt.Sprintf` yang berisi string "data" dan angka yang dipanggil dalam variabel `i`. Variabel `data` kemudian dikirim ke fungsi `doPrint`.

4. Di dalam fungsi `doPrint`, data diterima sebagai parameter message, lalu di-print menggunakan `fmt.Println`. Setelah itu, pemanggilan `wg.Done()` dijalankan untuk memberitahu bahwa tugas telah selesai.

5. Ketika fungsi `doPrint` dipanggil, program menambahkan 1 angka pada variabel `wg` menggunakan `wg.Add(1)`. Kemudian, fungsi `doPrint` dijalankan pada goroutine yang terpisah dengan menggunakan `go doPrint(&wg, data)`.

Setelah semua goroutine selesai, program menunggu dengan menggunakan `wg.Wait()` sampai semua fungsi yang dijalankan selesai sebelum program berakhir.

Jadi, program ini membuat beberapa goroutine yang mencetak pesan secara paralel, dan menunggu semua goroutine selesai sebelum program berakhir.
<br><br><br>


<b>Referensi:</b><br>
1. https://www.geeksforgeeks.org/using-waitgroup-in-golang/
2. https://linuxhint.com/golang-waitgroup/