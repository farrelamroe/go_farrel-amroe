package main

import (
	"fmt"
)

func power(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

func main() {
	var base, exponent int
	fmt.Printf("Masukkan basis (integer): ")
	fmt.Scanf("%d", &base)
	fmt.Print("Masukkan eksponen (integer): ")
	fmt.Scanf("%d", &exponent)
	fmt.Println("Hasilnya adalah", power(base, exponent))
}
