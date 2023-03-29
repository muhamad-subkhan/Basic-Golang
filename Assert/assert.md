<b>Assert</b><br>
Salah satu tools standar untuk testing yang bisa kita gunakan dari package testify.


Contoh

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, 30, 30, "perhitungan volume salah")
}
```

Pada kode di atas, terdapat fungsi `TestHitungVolume` yang akan dijalankan oleh Go test runner saat menjalankan unit test. Fungsi ini akan membandingkan nilai ekspektasi dengan nilai aktual dari perhitungan volume, menggunakan assert.Equal dari library testify/assert.

Dalam hal ini, assertion dilakukan dengan membandingkan dua parameter pertama yang diberikan, yaitu 20 dan 30. Jika nilai yang dibandingkan tidak sama, maka fungsi assert.Equal akan mengembalikan sebuah error dengan pesan "perhitungan volume salah" sebagai parameter ketiga.

Namun, karena dalam kode tersebut nilai ekspektasi dan aktual tidak sama, maka unit test akan gagal dan menunjukkan pesan error bahwa "perhitungan volume salah" pada baris kode tersebut. Seharusnya, nilai ekspektasi adalah 20 dan nilai aktual dari perhitungan volume harus dibandingkan dengan nilai tersebut agar unit test dapat berhasil.

<b>Referensi</b><br>
1. https://www.guru99.com/unit-testing-guide.html
2. https://www.youtube.com/watch?v=t9QJPE5vwhs&ab_channel=ProgrammerZamanNow
3. https://medium.com/yemeksepeti-teknoloji/how-to-write-unit-test-in-go-1df2b98ad510
4. https://medium.com/tunaiku-tech/unit-test-in-golang-57a2a896d90d



