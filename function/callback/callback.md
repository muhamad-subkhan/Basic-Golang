<b>Function Callback</b><br>
kemampuan suatu fungsi untuk menerima fungsi lain sebagai parameter dan kemudian memanggil fungsi tersebut di dalamnya. Fungsi yang diteruskan sebagai parameter ini disebut dengan callback function.


Contoh penggunaan function callback pada Go adalah sebagai berikut:

```go
package main

import "fmt"

func main() {
    // membuat fungsi calculate dengan parameter fungsi callback
    calculate(5, 7, func(result int) {
        fmt.Printf("Hasil penjumlahan adalah %d\n", result)
    })
}

func calculate(a, b int, callback func(int)) {
    // melakukan operasi perhitungan
    result := a + b
    
    // memanggil fungsi callback dengan parameter result
    callback(result)
}
```


Pada contoh di atas, kita membuat fungsi `calculate` yang menerima dua parameter bilangan `a` dan `b`, serta sebuah fungsi callback yang menerima satu parameter bertipe `int`. Fungsi `calculate` kemudian melakukan operasi perhitungan penjumlahan `a` dan `b`, dan mengirimkan hasilnya sebagai parameter ke dalam fungsi callback.

Di dalam fungsi `main`, kita memanggil fungsi `calculate` dengan memberikan dua bilangan dan sebuah fungsi callback yang akan dicetak hasilnya ke konsol menggunakan fungsi `Printf` dari package "fmt". Fungsi callback tersebut didefinisikan sebagai anonymous function.

Dengan menggunakan function callback pada Go, kita dapat membuat kode yang lebih modular dan reusable, serta dapat memisahkan antara operasi utama dengan aksi-aksi tambahan yang mungkin perlu dilakukan setelah operasi utama selesai. Function callback juga dapat digunakan untuk menangani event dan operasi yang membutuhkan waktu yang lama untuk dijalankan.


<b>Referensi</b><br>
1. https://www.marcelbelmont.com/golang-workshop/docs/lessons/callbacks.html