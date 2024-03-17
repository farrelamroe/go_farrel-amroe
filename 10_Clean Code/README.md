# Clean Code

1. Clean code adalah istilah code yang mudah dipahami, dibaca, dan diubah oleh prgrammer
2. Karakteristik clean code:
    - mudah dipahami
    - mudah dibaca dan dieja
    - singkat namun mendeskripsikan konteks
    - konsisten
    - hindari penambahan konteks yang tidak perlu
    - komen seperlunya, jangan tiap baris
    - membuat function jangan terlalu banyak argumen
    - gunakan konvensi
    - melakukan formatting, menggunakan prettier, memperhatikan indentasil, satu class maksimal 300 baris
3. Hindari membuat 1 fungsi membuat A, modifikasi B, mengecek C sekaligus
4. Buat fungsi yang bisa digunakan berulang-ulang untuk menghindari duplikasi
5. Refactoring adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal