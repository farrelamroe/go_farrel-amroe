package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for idx := 2; idx*idx <= number; idx++ {
		if number%idx == 0 {
			return false
		}
	}
	return true
}


func generatePrimes(n int, prima chan<- int) {
	for idx := 2; idx <= n; idx++ {
		if isPrime(idx) {
			prima <- idx 
		}
	}
	close(prima)
}

func powerPrimes(prima <-chan int, res chan<- int) {
	for prime := range prima {
		res <- prime * prime
	}
	close(res)
}

func main() {
	prima := make(chan int)
	res := make(chan int)


	go generatePrimes(100, prima)

	go powerPrimes(prima, res)


	fmt.Println("Hasil perpangkatan bilangan prima:")
	for result := range res {
		fmt.Println(result)
	}
}
