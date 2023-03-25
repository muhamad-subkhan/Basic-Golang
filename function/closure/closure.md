<b>Fuction Closure</b><br>
kemampuan suatu fungsi untuk mengakses dan memanipulasi variabel-variabel yang berada di luar ruang lingkup fungsi tersebut (outer variable). Dalam Go, function closure dapat diimplementasikan dengan menggunakan anonymous function atau function literal.

Ketika sebuah anonymous function di-deklarasikan di dalam suatu fungsi, ia dapat mengakses dan memanipulasi variabel-variabel yang berada di luar ruang lingkup fungsi tersebut. Variabel-variabel ini biasanya disebut dengan captured variables atau free variables. Anonymous function tersebut dapat digunakan untuk membuat closure, yang merupakan kumpulan dari anonymous function beserta captured variables-nya.

Contoh penggunaan function closure pada Go adalah sebagai berikut:

```go 
package main


import "fmt"

func main() {
    // membuat closure yang mengambil variabel "counter" dari luar ruang lingkup
    counter := 0
    increment := func() {
        counter++
        fmt.Println(counter)
    }
    
    // menggunakan closure untuk memanggil anonymous function
    increment() // Output: 1
    increment() // Output: 2
    increment() // Output: 3
}
```

Pada contoh di atas, kita membuat closure dengan anonymous function yang mengambil variabel `counter` dari luar ruang lingkup fungsi `main`. Variabel `counter` awalnya diinisialisasi dengan nilai `0`.

Ketika kita memanggil closure tersebut dengan `increment()`, nilai `counter` akan di-increment dan dicetak ke konsol menggunakan fungsi `Println` dari package "fmt". Kita dapat memanggil closure ini berkali-kali, dan setiap kali nilai counter akan di-increment dan dicetak ke konsol.

Dalam hal ini, closure pada Go memungkinkan kita untuk membuat fungsi yang lebih fleksibel dan dapat mengakses variabel-variabel yang berada di luar ruang lingkup fungsi tersebut. Closure juga dapat digunakan untuk mengimplementasikan konsep-konsep seperti factory function, iterator, dan lain sebagainya.


<b>Referensi</b><br>
1. https://www.programiz.com/golang/closure