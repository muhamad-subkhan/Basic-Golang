<b>Struct</b><br>
sebuah tipe data terstruktur yang digunakan untuk merepresentasikan data dalam bentuk gabungan dari beberapa tipe data yang berbeda. Struct terdiri dari beberapa field, dimana masing-masing field memiliki tipe data yang berbeda.



<b>Contoh</b>

```go
package main

import "fmt"

func main() {

	type identity struct {
		name string
		age  int
	}
	identitas  := identity{name: "John doe", age: 25}

	// fmt.Println(identitas)
	fmt.Println(identitas.name)
	fmt.Println(identitas.age)
}
```
<b>name, age</b> -> Properties
<b>John doe, 25</b> -> value/nilai


<b>Referensi</b>
1. https://gobyexample.com/structs
2. https://www.youtube.com/watch?v=NMTN543WVQY&ab_channel=HiteshChoudhary
