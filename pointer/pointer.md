<b>pointer</b><br>
variabel yang menyimpan alamat memori dari variabel lainnya.

```go
package main

import "fmt"

func main() {

	var a int = 42
	var p *int = &a

	fmt.Println("Nilai a: ", a)
	fmt.Println("Alamat a: ", &a)
	fmt.Println("Nilai p: ", p)
	fmt.Println("Nilai yang ditunjuk oleh p: ", *p)

	*p = 21
	fmt.Println("Nilai a setelah diubah melalui p: ", a)
}

```

<u>
<b>referensi</b>
</u>

1. https://golangbyexample.com/all-data-types-in-golang-with-examples/#Pointers
