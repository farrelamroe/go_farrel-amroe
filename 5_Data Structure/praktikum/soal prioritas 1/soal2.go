package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
func main() {
	res := 0
	var n int
	fmt.Printf("Masukkan panjang list: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if primeNumber(arr[i]) == true {
			res += arr[i]
		}
	}
	fmt.Println(res)
}
