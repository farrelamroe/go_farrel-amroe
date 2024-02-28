package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func separatePrimes(numbers []int) ([]int, []int) {
	primes := []int{}
	others := []int{}
	for _, num := range numbers {
		if isPrime(num) {
			primes = append(primes, num)
		} else {
			others = append(others, num)
		}
	}
	return primes, others
}

func main() {
	var n int
	fmt.Printf("Masukkan jumlah bilangan: ")
	fmt.Scanln(&n)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	primes, others := separatePrimes(numbers)
	fmt.Println("Input:", numbers)
	fmt.Println("Output:", primes, others)
}
