package main

import (
	"fmt"
)

var memo map[int]int

func fibonacci(number int) int {
	if val, ok := memo[number]; ok {
		return val
	}

	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}

	fib := fibonacci(number-1) + fibonacci(number-2)

	memo[number] = fib

	return fib
}

func main() {

	memo = make(map[int]int)

	var number int
	fmt.Print("Masukkan angka untuk menghitung Fibonacci: ")
	fmt.Scan(&number)

	fmt.Printf("Angka Fibonacci ke-%d adalah: %d\n", number, fibonacci(number))
}
