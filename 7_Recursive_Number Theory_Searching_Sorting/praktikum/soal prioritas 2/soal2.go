package main

import (
	"fmt"
	"math"
)

func primeFactors(n int) {
	factors := []int{}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	for i, factor := range factors {
		fmt.Print(factor)
		if i < len(factors)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func main() {
	var input int
	fmt.Printf("Masukkan bilangan: ")
	fmt.Scanln(&input)
	primeFactors(input)
}
