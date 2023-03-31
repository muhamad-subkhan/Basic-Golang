<b>Race Condition</b> adalah kondisi dimana lebih dari satu goroutine, mengakses data yang sama pada waktu bersamaan (benar-benar bersamaan). Ketika hal itu terjadi, nilai data tersebut akan menjadi kacau. Dalam concurrency programming situasi seperti ini sering tejadi.


<b>Mutex</b><br>
digunakan untuk mengendalikan akses ke sebuah resource yang bersifat shared (dibagi) oleh beberapa goroutine. Dalam penggunaan mutex, goroutine yang ingin mengakses resource tersebut harus memperoleh akses eksklusif terlebih dahulu dengan cara memasuki area kritis (critical section) dan keluar dari area kritis tersebut setelah selesai menggunakan resource.

Pada Go, mutex diimplementasikan melalui tipe data `sync.Mutex` yang menyediakan dua metode utama, yaitu `Lock()` dan `Unlock()`. `Lock()` digunakan untuk memperoleh akses eksklusif ke resource, sedangkan `Unlock()` digunakan untuk melepaskan akses eksklusif tersebut setelah selesai menggunakan resource.
<br><br><br>

Contoh:

```go
package main

import (
    "fmt"
    "sync"
)

var count int
var mu sync.Mutex

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
    fmt.Println("Count incremented to:", count)
}

func main() {
    for i := 0; i < 10; i++ {
        go increment()
    }
    fmt.Scanln()
}

```

Dalam contoh di atas, kita menggunakan mutex untuk mengendalikan akses ke variabel `count` yang bersifat shared oleh beberapa goroutine. Fungsi `increment()` digunakan untuk menambahkan nilai variabel `count` dengan menggunakan mekanisme mutex.

Pada saat goroutine ingin mengakses resource shared yang sama, ia harus memperoleh akses eksklusif dengan cara memanggil metode `Lock()` pada mutex terlebih dahulu. Setelah selesai menggunakan resource, goroutine harus memanggil metode `Unlock()` untuk melepaskan akses eksklusif tersebut dan memberikan kesempatan pada goroutine lain untuk mengakses resource.

Dalam contoh di atas, fungsi `increment()` menggunakan mekanisme `defer` untuk memastikan bahwa metode `Unlock()` selalu dipanggil, bahkan jika terjadi panic pada saat fungsi sedang dijalankan. Hal ini penting untuk memastikan bahwa akses eksklusif terhadap resource selalu dilepaskan setelah selesai digunakan.


<b>Referensi</b><br>
1. https://www.sohamkamani.com/golang/mutex/
2. https://kodingin.com/golang-mutex/
3. https://www.geeksforgeeks.org/mutex-in-golang-with-examples/