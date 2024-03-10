# Concurrent Programming

1. Pada program sequential, harus menyelesaikan task sebelumnya lalu lanjut ke task baru
2. Pada program parallel, berbagai task bisa diselesaikan secara bersamaan
3. Para program concurrent, Beberapa tugas dapat dijalankan secara independen dan mungkin muncul secara simultan
4. Concurrency pada Go membuat mudah dalam membuat program parallel yang memanfaatkan mesin dengan multiple processor
5. Goroutine adalah fungsi atau metode yang berjalan bersamaan (independen) dengan fungsi atau metode lain
6. Gomaxprocs digunakan untuk kontrol jumlah operasi systemthreads
7. Channel adalah objek komunikasi yang dengannya goroutine dapat berkomunikasi satu sama lain
8. Select membuatnya lebih mudah untuk mengontrol komunikasi data melalui satu atau banyak saluran
9. Race condition terjadi dimana 2 threads mengakses memori pada waktu bersamaan. salah satunya adalah writing