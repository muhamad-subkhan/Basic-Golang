<b>Select Channel</b><br>
Select channel adalah salah satu fitur pada bahasa pemrograman Go (Golang) yang memungkinkan pengguna untuk memantau dan menangani beberapa channel secara asynchronous dalam sebuah blok pernyataan tunggal. Fitur ini berguna untuk menyelesaikan masalah yang terkait dengan pembacaan atau penulisan data pada beberapa channel secara bersamaan tanpa harus menggunakan beberapa goroutine.

Dalam select channel, kita dapat memantau beberapa channel secara bersamaan menggunakan sintaks select dan menentukan aksi apa yang akan dilakukan jika salah satu channel tersebut memiliki data yang siap untuk dibaca atau ditulis.<br><br><br>

<img src='https://res.cloudinary.com/practicaldev/image/fetch/s--aZu0BCUa--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://res.cloudinary.com/ahmedash/image/upload/v1575824142/dev.to/worker-queue.png' alt='select channel' />
<br><br><br>

Contoh:

```go
package main

import "fmt"

func main() {
    ch1 := make(chan int)
    ch2 := make(chan string)

    go func() {
        ch1 <- 10
    }()

    go func() {
        ch2 <- "Hello, World!"
    }()

    for i := 0; i < 2; i++ {
        select {
        case val := <-ch1:
            fmt.Println("Received value from ch1:", val)
        case val := <-ch2:
            fmt.Println("Received value from ch2:", val)
        }
    }
}

```


Dalam contoh di atas, kita membuat dua channel, yaitu ch1 dan ch2. Selanjutnya, kita membuat dua goroutine yang masing-masing mengirimkan data ke channel ch1 dan ch2 secara asynchronous.

Selanjutnya, kita menggunakan sintaks select pada loop utama untuk memantau kedua channel tersebut. Jika ada data yang siap untuk dibaca pada salah satu channel, maka aksi yang sesuai akan diambil.

Pada contoh di atas, kita memeriksa kedua channel secara bergantian menggunakan loop dengan kondisi i < 2. Jika ada data yang siap dibaca pada ch1, maka nilai tersebut akan dicetak pada konsol dengan pesan "Received value from ch1". Begitu juga jika data yang siap dibaca berasal dari ch2, maka pesan "Received value from ch2" akan dicetak pada konsol.
<br><br><br>


<b>Referensi</b><br>
1. https://www.geeksforgeeks.org/select-statement-in-go-language/
2. https://www.youtube.com/watch?v=1c7ttSJDMAI