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