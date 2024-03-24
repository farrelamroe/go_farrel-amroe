# System Design

1. Diagram Design Tools: 
    - smartdraw
    - lucidchart
    - whimsical
    - draw.io
    - visio
2. Common used software design:
    - flowchart
    - use case
    - ERD
    - HLA
3. Key Characteristics of Distributed Systems
    - Scalability: kemampuan sistem, proses, atau jaringan untuk tumbuh dan mengelola peningkatan permintaan.
    - Reliability: probabilitas sistem akan gagal dalam periode tertentu.
    - Availability: waktu suatu sistem tetap beroperasi untuk melakukan fungsi yang diperlukan dalam periode tertentu.
    - Efficiency: sumsikan kita memiliki operasi yang berjalan secara terdistribusi dan memberikan satu set item sebagai hasilnya. Dua ukuran standar efisiensinya adalah waktu respons (atau latensi) yang menunjukkan penundaan untuk mendapatkan item pertama dan throughput (atau bandwidth) yang menunjukkan jumlah item yang dikirim dalam unit waktu tertentu (misalnya, satu detik).
    - Serviceability or Manageability: kesederhanaan dan kecepatan dimana suatu sistem dapat diperbaiki atau dipelihara.
4. Work Queue adalah framework untuk membangun aplikasi master-worker besar yang menjangkau ribuan mesin yang diambil dari cluster, cloud, dan grid.
5. Load Balancer (LB) adalah komponen penting lainnya dari setiap sistem terdistribusi. Ini membantu menyebarkan lalu lintas di sekelompok server untuk meningkatkan respons dan ketersediaan aplikasi, situs web, atau database.
6. Monolithic Application memiliki basis kode tunggal dengan beberapa modul. Modul dibagi sebagai fitur bisnis atau fitur teknis. Ini memiliki sistem build tunggal yang membangun seluruh aplikasi dan / atau ketergantungan. Ini juga memiliki biner tunggal yang dapat dieksekusi atau digunakan.
7. Microservice adalah layanan yang dapat digunakan secara independen yang dimodelkan di sekitar domain bisnis. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur menawarkan banyak pilihan untuk memecahkan masalah yang mungkin Anda hadapi. Oleh karena itu, arsitektur layanan mikro didasarkan pada beberapa layanan mikro yang berkolaborasi.
8. SQL vs NoSQL
    - Database relasional terstruktur dan memiliki skema yang telah ditentukan seperti buku telepon yang menyimpan nomor telepon dan alamat.
    - Database non-relasional tidak terstruktur, dan memiliki skema dinamis seperti folder file yang menyimpan segala sesuatu mulai dari alamat dan nomor telepon seseorang hingga 'suka' Facebook dan preferensi belanja online mereka.
9. Cache yang digunakan dalam data yang baru-baru ini diminta kemungkinan akan diminta lagi. Mereka digunakan di hampir setiap lapisan komputasi: perangkat keras, sistem operasi, browser web, aplikasi web, dan banyak lagi. Cache seperti memori jangka pendek: ia memiliki jumlah ruang terbatas, tetapi biasanya lebih cepat daripada sumber data asli dan berisi item yang paling baru diakses. Cache dapat ada di semua tingkatan dalam arsitektur, tetapi sering ditemukan pada tingkat terdekat dengan ujung depan di mana mereka diimplementasikan untuk mengembalikan data dengan cepat tanpa membebani tingkat hilir.
10. Redundancy adalah duplikasi komponen atau fungsi penting dari suatu sistem dengan maksud meningkatkan keandalan sistem, biasanya dalam bentuk cadangan atau gagal-aman, atau untuk meningkatkan kinerja sistem yang sebenarnya.