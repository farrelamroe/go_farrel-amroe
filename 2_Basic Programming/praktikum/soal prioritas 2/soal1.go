package main

import "fmt"

func main() {
	var nilai int
	fmt.Printf("Masukkan nilai: ")
	fmt.Scanf("%d", &nilai)

	for i := 1; i <= nilai; i++ {
		if nilai%i == 0 {
			if i%2 == 0 {
				fmt.Printf("I")
			} else {
				fmt.Printf("O")
			}
		}
	}
}
