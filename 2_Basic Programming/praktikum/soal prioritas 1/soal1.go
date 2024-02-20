package main

import "fmt"

func main() {
	pi := 3.14
	var r, t float64
	fmt.Printf("Masukkan nilai jari jari: ")
	fmt.Scanf("%f", &r)
	fmt.Printf("Masukkan nilai tinggi: ")
	fmt.Scanf("%f", &t)
	v := pi * r * r * t

	fmt.Println("Volume tabung dengan jari jari", r, "dan tinggi", t, "adalah", v)

}
