<b>Context</b><br>
digunakan untuk mengelola waktu hidup (lifecycle) dan pembatalan operasi di lingkungan multigoroutine. Context memberikan mekanisme yang aman dan sederhana untuk membagikan nilai-nilai (value) antar goroutine, seperti koneksi database atau token otorisasi.

Context memiliki fungsi dan kegunaan utama sebagai berikut:

Memungkinkan pembatalan operasi: Context menyediakan kemampuan untuk membatalkan operasi yang sedang berjalan di goroutine, baik itu operasi input/output (I/O), operasi jaringan, atau operasi lainnya. Hal ini berguna ketika kita ingin menghentikan operasi yang tidak relevan atau sudah tidak perlu lagi.

Mengendalikan waktu hidup: Context memungkinkan pengguna untuk menentukan waktu hidup (lifecycle) sebuah goroutine, membatasi durasi waktu eksekusi sebuah operasi, atau memberikan batasan waktu (timeout) untuk sebuah operasi.

Menyimpan data pada konteks: Context dapat menyimpan data dengan kunci (key) pada setiap goroutine. Data ini dapat digunakan untuk berbagi informasi antar goroutine.

Beroperasi di lingkungan multigoroutine: Context memungkinkan komunikasi antar goroutine tanpa menggunakan mekanisme yang rentan deadlock dan race condition.

Dalam pengembangan aplikasi Go, Context menjadi sangat penting ketika aplikasi menangani beberapa permintaan sekaligus. Context memberikan mekanisme untuk mengelola waktu hidup permintaan dan menghindari penggunaan memori yang berlebihan. Oleh karena itu, penggunaan Context sangat disarankan dalam pengembangan aplikasi Go yang handal dan aman.
<br>
<img src="https://res.cloudinary.com/practicaldev/image/fetch/s--nbKyLWT1--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/6eos6t546i539bbwr0rl.png" alt="context" />
<br><br>

<b>Contoh Program</b><br>

```go

``










<br><br><br>
<b>Referensi:</b><br>
1. https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
2. https://www.youtube.com/watch?v=LSzR0VEraWw&ab_channel=justforfunc%3AProgramminginGo
3. https://www.youtube.com/watch?v=h2RdcrMLQAo&feature=youtu.be&ab_channel=TutorialEdge