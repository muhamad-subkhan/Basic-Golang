<b>Context</b><br>
digunakan untuk mengelola waktu hidup (lifecycle) dan pembatalan operasi di lingkungan multigoroutine. Context memberikan mekanisme yang aman dan sederhana untuk membagikan nilai-nilai (value) antar goroutine, seperti koneksi database atau token otorisasi.

Context memiliki fungsi dan kegunaan utama sebagai berikut:

Memungkinkan pembatalan operasi: Context menyediakan kemampuan untuk membatalkan operasi yang sedang berjalan di goroutine, baik itu operasi input/output (I/O), operasi jaringan, atau operasi lainnya. Hal ini berguna ketika kita ingin menghentikan operasi yang tidak relevan atau sudah tidak perlu lagi.

Mengendalikan waktu hidup: Context memungkinkan pengguna untuk menentukan waktu hidup (lifecycle) sebuah goroutine, membatasi durasi waktu eksekusi sebuah operasi, atau memberikan batasan waktu (timeout) untuk sebuah operasi.

Menyimpan data pada konteks: Context dapat menyimpan data dengan kunci (key) pada setiap goroutine. Data ini dapat digunakan untuk berbagi informasi antar goroutine.

Beroperasi di lingkungan multigoroutine: Context memungkinkan komunikasi antar goroutine tanpa menggunakan mekanisme yang rentan deadlock dan race condition.

Dalam pengembangan aplikasi Go, Context menjadi sangat penting ketika aplikasi menangani beberapa permintaan sekaligus. Context memberikan mekanisme untuk mengelola waktu hidup permintaan dan menghindari penggunaan memori yang berlebihan. Oleh karena itu, penggunaan Context sangat disarankan dalam pengembangan aplikasi Go yang handal dan aman.
<br><br><br>
<img src="https://res.cloudinary.com/practicaldev/image/fetch/s--nbKyLWT1--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/6eos6t546i539bbwr0rl.png" alt="context" />
<br><br><br>

<b>Contoh Program</b><br>

```go
package main

import (
	"context"
	"fmt"
)

func main() {

	contextParent := context.Background()

	ctxSatu := context.WithValue(contextParent, "key1", "Hello World")
	ctxDua  := context.WithValue(ctxSatu, "key2", "Hello girls")
	ctxTiga := context.WithValue(ctxDua, "key3", "Hello boys")
	ctxEmpat := context.WithValue(contextParent, "key4", "Hello Children")
	ctxLima := context.WithValue(ctxSatu, "key5", "Hello Later")

	fmt.Println(ctxLima.Value("key5"))
	fmt.Println(ctxLima.Value("key4"))
	fmt.Println(ctxLima.Value("key3"))
	fmt.Println(ctxLima.Value("key2"))
	fmt.Println(ctxLima.Value("key1"))
	fmt.Println(ctxTiga.Value("key1"))
	fmt.Println(ctxEmpat.Value("key1"))
}
``

Pada program ini, terdapat fungsi "main" yang melakukan beberapa operasi dengan menggunakan context. Pertama-tama, dibuat context parent dengan menggunakan fungsi "context.Background()". Context parent ini kemudian digunakan sebagai parameter pada beberapa fungsi "context.WithValue()", yang akan mengembalikan sebuah context baru dengan nilai tertentu yang ditentukan.

Context baru tersebut diberi nama "ctxSatu", "ctxDua", "ctxTiga", "ctxEmpat", dan "ctxLima". Pada setiap context baru tersebut, dilakukan pemberian nilai pada beberapa key tertentu.

Setelah context baru dibuat, program mencetak nilai dari masing-masing key pada beberapa context yang berbeda menggunakan fungsi "Value()" pada setiap context tersebut. Output program menampilkan nilai-nilai pada setiap key di dalam setiap context yang telah dibuat.








<br><br><br>
<b>Referensi:</b><br>
1. https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
2. https://www.youtube.com/watch?v=LSzR0VEraWw&ab_channel=justforfunc%3AProgramminginGo
3. https://www.youtube.com/watch?v=h2RdcrMLQAo&feature=youtu.be&ab_channel=TutorialEdge