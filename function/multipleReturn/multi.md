<b>Function Multiple Returns</b><br>
kemampuan suatu fungsi untuk mengembalikan lebih dari satu nilai. Dalam bahasa Go, sebuah fungsi dapat mengembalikan nilai lebih dari satu, dengan memisahkan nilai-nilai yang dikembalikan dengan tanda koma.

<b>Contoh</b><br>

```go
package main


import "fmt"


func main() {
	add, sub := calculate(10, 5)
    fmt.Println("Addition result:", add)
    fmt.Println("Subtraction result:", sub)
}


func calculate(a int, b int) (int, int) {
	return a + b, a - b
}

```

Pada contoh di atas, fungsi calculate mengambil dua parameter a dan b, kemudian mengembalikan dua nilai, yaitu a + b dan a - b. Di dalam fungsi main, fungsi calculate dipanggil dengan memberikan nilai 10 dan 5 sebagai parameter.

Kemudian, hasil kembalian dari fungsi calculate yang terdiri dari hasil penjumlahan dan hasil pengurangan dari kedua parameter disimpan dalam variabel add dan sub menggunakan teknik multiple assignment.


<b><i>gampangnya</i></b><br> ketika sebuah function memiliki nilai return lebih 1, maka harus dimasukan kedalan ( )