package main

import (
	"fmt"
)

func rotateSlice(slice []int) {
	length := len(slice)
	if length <= 1 {
		return
	}
	mid := length / 2
	for i := 1; i < mid; i++ {
		slice[length-i], slice[i] = slice[i], slice[length-i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan panjang slice: ")
	fmt.Scan(&n)

	slice := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Println("Input:", slice)
	rotateSlice(slice)
	fmt.Println("Output:", slice)
}
