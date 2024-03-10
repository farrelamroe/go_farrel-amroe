package main

import (
	"fmt"
)

func isComposite(number int) bool {
	if number <= 1 {
		return false
	}
	for idx := 2; idx*idx <= number; idx++ {
		if number%idx == 0 {
			return true
		}
	}
	return false
}

func generateComposites(n int, composites chan<- int) {
	for idx := 2; idx <= n; idx++ {
		if isComposite(idx) {
			composites <- idx
		}
	}
	close(composites)
}

func processComposites(composites <-chan int, res chan<- string) {
	for composite := range composites {
		powerPrimes := composite * composite
		if powerPrimes%2 == 0 {
			res <- "Ping"
		} else {
			res <- "Pong"
		}
	}
	close(res)
}

func main() {
	composites := make(chan int)
	res := make(chan string)

	go generateComposites(100, composites)

	go processComposites(composites, res)

	fmt.Println("Hasil:")
	for result := range res {
		fmt.Println(result)
	}
}
