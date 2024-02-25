package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menentukan apakah sebuah bilangan prima atau bukan
func primeNumber(n int) string {
	trueSentence := "Bilangan prima"
	falseSentence := "Bukan bilangan prima"
	if n <= 1 {
		return falseSentence
	}

	if n == 2 {
		return trueSentence
	}

	if n%2 == 0 {
		return falseSentence
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return falseSentence
		}
	}

	return trueSentence
}

func main() {
	var number int
	fmt.Printf("Masukkan bilangan: ")
	fmt.Scanf("%d", &number)
	fmt.Println("Nilai input adalah", primeNumber(number))
}
