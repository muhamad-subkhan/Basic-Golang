<b>Make </b><br>
sebuah fungsi built-in yang digunakan untuk membuat slice, map, dan channel yang dikelola oleh runtime. Fungsi make mengalokasikan dan menginisialisasi data pada slice, map, atau channel yang dibuat, dan mengembalikan nilai bertipe pointer dari tipe data yang sesuai.


<b>Contoh</b>

```go
package main

import "fmt"

func main() {

	var fruits = make([]string, 2)
	fruits[0] = "foo"
	fruits[1] = "bar"

	fmt.Println(fruits)
}
```

penggunaan fungsi make untuk membuat sebuah slice fruits dengan panjang awal 2 

Setelah membuat slice, program kemudian mengisi setiap elemen slice dengan nilai "foo" dan "bar" pada indeks 0 dan 1.


<b>Referensi</b><br>
1. https://www.golangprograms.com/how-to-create-slice-using-make-function-in-golang.html
2. https://www.golangprograms.com/golang-package-examples/how-to-create-map-using-the-make-function-in-go.html
3. https://www.programiz.com/golang/channel#channel