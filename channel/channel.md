<b?>Channel</b><br>
digunakan untuk memfasilitasi komunikasi dan sinkronisasi antara Goroutine (concurant thread of execution). Channel menyediakan sebuah mekanisme dimana satu goroutine dapat mengirimkan data ke goroutine lain atau menerima data goroutine lain dengan cara yang aman dan sinkronis.
<br><br><br>

<image src="./Untitled-Diagram46.jpg" />
<br><br><br>

Contoh Code:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)
	sayHello := func (who string) {
		var data = fmt.Sprintf("Hallo %s", who)
		message <- data;
	}


	go sayHello("John Doe")
	go sayHello("Jhonny")
	go sayHello("kitty")
	
	messageSatu := <- message
	fmt.Println(messageSatu)

	messageDua := <- message
	fmt.Println(messageDua)

	messageTiga := <- message
	fmt.Println(messageTiga)

}
```

Code di atas merupakan contoh penggunaan channel pada Go untuk melakukan komunikasi antara beberapa goroutine. Pertama-tama, kode tersebut mengatur jumlah maksimum prosesor yang dapat digunakan oleh aplikasi menggunakan fungsi `runtime.GOMAXPROCS(2)`.

Kemudian, kode membuat sebuah channel dengan tipe data `string` menggunakan fungsi `make(chan string)`. Channel ini digunakan untuk mengirimkan pesan dari goroutine yang satu ke goroutine yang lain.

Selanjutnya, kode membuat tiga goroutine yang masing-masing menjalankan fungsi `sayHello` dengan parameter yang berbeda. Fungsi `sayHello` akan membuat pesan dengan format "Hallo {nama}" dan mengirimkannya melalui channel message.

Setelah itu, kode menggunakan operator `<-` untuk menerima pesan dari channel `message`. Dalam contoh ini, kode menerima pesan sebanyak tiga kali menggunakan operator `<-` secara berurutan, dan menampilkan pesan tersebut ke layar dengan fungsi `fmt.Println()`.

Dalam keseluruhan, kode tersebut merupakan contoh sederhana penggunaan channel dan goroutine pada Go untuk melakukan komunikasi dan sinkronisasi antara beberapa proses secara concurrent.
<br><br><br>
<b>Output</b><br>
```go
Hallo kitty
Hallo John Doe
Hallo Jhonny
```
<br><br><br>
Kenapa Hallo kitty menjadi output pertama? Hal itu terjadi karena dalam contoh kode tersebut, terdapat tiga goroutine yang berjalan secara concurrent untuk menjalankan fungsi `sayHello` yang masing-masing mengirimkan pesan ke channel `message`. Urutan eksekusi goroutine ini tidak dapat diprediksi, karena sistem operasi yang menjalankan program akan menentukan kapan setiap goroutine berjalan dan kapan goroutine akan berhenti sementara untuk memberikan kesempatan pada goroutine lain untuk berjalan.

Dalam contoh kode tersebut, meskipun pesan "Hallo kitty" ditulis terakhir dalam kode, tetapi pesan tersebut mungkin saja dikirimkan pertama kali oleh goroutine yang menjalankan fungsi sayHello("kitty"). Karena goroutine-goroutine ini berjalan secara concurrent, maka tidak dapat diprediksi dengan pasti urutan pesan mana yang akan dikirimkan lebih dulu dan mana yang akan dikirimkan kemudian.


<b>Referensi</b><br>
1. https://golangbot.com/channels/
2. https://www.youtube.com/watch?v=e4bu9g-bYtg&ab_channel=TutorialEdge
3. https://www.geeksforgeeks.org/channel-in-golang/
4. https://www.youtube.com/watch?v=LgCmPHqAuf4&ab_channel=GolangDojo
