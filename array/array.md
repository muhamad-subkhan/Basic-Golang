<b>Array</b><br>
kumpulan data yang akan disimpan dalam variable.

contoh 

```go
package main

import "fmt"

func main() {

	fruits := [5]string{"apple", "banana", "orange", "melon", "cherry"}
	fmt.Println("\nJumlah Fruits:", len(fruits))
	fmt.Println("\nisi fruits:",fruits)
}
```

<b>Fruits</b> adalah variable yang memiliki isian array dengan 5 value <br><br>
jumlah isian array disebut dengan index. dalam contoh diatas, fruits memiliki index 5 yaitu 0, 1, 2, 3, 4 <br> dalam indexing angka pertama ialah 0 bukan 1. maka dari itu ketika fruits memiliki 5 value artinya <b><i>apple -> index ke 0, banana -> index ke 1 dan seterusnya</i></b> <br>


<b>Referensi</b>
1. https://www.youtube.com/watch?v=YS4e4q9oBaU&t=6473s&ab_channel=freeCodeCamp.org
