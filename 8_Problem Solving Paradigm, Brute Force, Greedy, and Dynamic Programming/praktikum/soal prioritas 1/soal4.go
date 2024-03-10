package main

import "fmt"

func SimpleEquations(a, b, c int) {
    // Mencari nilai x, y, dan z dari 1 hingga 10000
    for x := 1; x <= 10000; x++ {
        for y := 1; y <= 10000; y++ {
            for z := 1; z <= 10000; z++ {
                // Memeriksa apakah nilai-nilai tersebut memenuhi hubungan
                if x+y+z == a && x*y+x*z+y*z == b && x*y*z == c {
                    fmt.Println(x, y, z)
                    return
                }
            }
        }
    }
    // Jika tidak ada solusi yang ditemukan
    fmt.Println("no solution")
}

func main() {
    // Meminta input dari pengguna
    var a, b, c int
    fmt.Println("Masukkan nilai A, B, dan C (pisahkan dengan spasi):")
    fmt.Scan(&a, &b, &c)

    // Memanggil fungsi SimpleEquations dengan input dari pengguna
    SimpleEquations(a, b, c)
}
