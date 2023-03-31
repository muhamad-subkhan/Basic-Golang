<b>Buffer</b><br>
Buffered channel adalah salah satu jenis channel pada bahasa pemrograman Go (Golang) yang memungkinkan pengiriman dan penerimaan data antar goroutine dengan mekanisme blocking dan non-blocking. Perbedaan utama buffered channel dengan unbuffered channel adalah buffered channel memiliki kapasitas tertentu yang dapat menampung beberapa nilai data dalam antrian sebelum channel tersebut dibaca oleh goroutine penerima.
<br><br><br>
<img src='https://www.ardanlabs.com/images/goinggo/Screen+Shot+2014-02-17+at+8.38.15+AM.png' atl="buffered-channel"/>
<br><br><br>



```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 3)
	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
```

<b><u>Penjelasan</u></b><br>
Pertama-tama, dilakukan pengaturan jumlah maksimum goroutine yang diizinkan untuk berjalan secara simultan menggunakan fungsi "runtime.GOMAXPROCS(2)", yang mengatur jumlah maksimum goroutine sebanyak 2.

Selanjutnya, dibuat channel bernama "messages" dengan kapasitas buffer sebesar 3 menggunakan fungsi "make(chan int, 3)".

Kemudian, dilakukan pembuatan goroutine dengan fungsi anonim yang berisi sebuah loop tak terbatas, yang akan terus melakukan pengambilan data dari channel "messages" menggunakan operator "<-" dan mencetak nilai data tersebut pada konsol. Goroutine tersebut dijalankan secara asynchronous dengan perintah "go func() {...}()".

Pada bagian utama program, dilakukan loop dengan variabel "i" dimulai dari 0 hingga 4, dan pada setiap iterasi loop dilakukan pengiriman data ke channel "messages" menggunakan operator "<-".

Ketika jumlah data yang dikirim sudah mencapai 3, maka goroutine yang menunggu untuk menerima data dari channel akan mulai bekerja. Data akan diambil dari buffer pada channel "messages" secara berurutan (FIFO) dan dicetak pada konsol.

Jika jumlah data yang dikirim lebih dari 3, maka goroutine akan menunggu sampai ada ruang kosong di dalam buffer sebelum menerima data berikutnya.

Setelah loop selesai, program akan selesai dieksekusi dan keluar dari program.
<br><br><br>
<b>Referensi</b><br>
1. https://www.educba.com/golang-buffer/
2. https://www.youtube.com/watch?v=NoDRq6Twkts&ab_channel=TutorialEdge