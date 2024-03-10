package main

import "fmt"

func fibonacci(n int) int {
	a, b := 0, 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	res := 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
		res += c
	}
	return res
}

func main() {
	var n int
	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}
