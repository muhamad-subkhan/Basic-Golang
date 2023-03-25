<b>Recover</b><br>
sebuah fungsi bawaan yang dapat digunakan untuk menangkap panic yang terjadi selama runtime program. Fungsi recover() digunakan untuk membatalkan panic dan mengembalikan kendali program ke program yang normal.

Ketika sebuah program Go mengalami panic, maka program tersebut akan berhenti secara tiba-tiba dan menampilkan informasi tentang kesalahan yang terjadi. Namun, jika kita menggunakan recover(), kita bisa menangani panic dan mengembalikan program ke keadaan yang normal.

Cara kerja recover() adalah dengan menempatkan fungsi tersebut di dalam sebuah blok defer. Ketika terjadi panic, program akan menjalankan semua blok defer di dalam fungsi tersebut, termasuk blok defer yang berisi recover(). Fungsi recover() akan mengambil nilai panic yang terjadi dan mengembalikan kendali program ke kondisi normal.

Perlu diingat bahwa recover() hanya akan berhasil menangkap panic jika dipanggil di dalam blok defer. Jika recover() dipanggil di luar blok defer atau tidak dijalankan pada saat terjadi panic, maka program akan tetap berhenti secara tiba-tiba dan menampilkan pesan error. Oleh karena itu, penggunaan recover() harus diperhatikan dengan baik agar tidak menimbulkan behavior yang tidak diinginkan.




<b>Referensi</b><br>
1. https://gobyexample.com/recover
2. https://www.educative.io/answers/what-is-recover-in-golang