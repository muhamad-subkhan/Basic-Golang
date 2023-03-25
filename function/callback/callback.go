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