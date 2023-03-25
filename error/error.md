<b>Error</b>⚠️<br>
 pesan yang muncul ketika program Go mengalami masalah dalam melakukan operasi. Ada beberapa jenis error yang dapat terjadi dalam program Go, seperti error tipe, error runtime, dan panic.

1. Error tipe: error tipe terjadi ketika tipe variabel yang digunakan tidak cocok dengan tipe yang diharapkan oleh program. Contoh: jika sebuah variabel bertipe string dioperasikan dengan operator yang seharusnya hanya digunakan untuk tipe numerik, maka program akan menghasilkan error tipe.

2. Error runtime: error runtime terjadi ketika program sedang dijalankan dan menemukan masalah, seperti pembagian dengan nol atau akses indeks array yang tidak valid.

3. Panic: panic terjadi ketika program mengalami masalah yang tidak bisa diatasi dan program menghentikan dirinya secara tiba-tiba.

Untuk menangani error di Go, kita bisa menggunakan mekanisme penanganan error seperti `if err != nil` atau dengan menggunakan `defer panic-recover`. Dengan mekanisme tersebut, kita bisa menangani error dan memberikan respons yang sesuai ketika program mengalami masalah.


<b>Referensi</b><br>
1. https://earthly.dev/blog/golang-errors/