package main

import "fmt"

func min(data ...int) int {
	res := data[0]
	for i := 0; i < len(data); i++ {
		if res < data[i] {
			res = data[i]
		}
	}
	return res
}

func main() {
	var panjang int
	fmt.Printf("Masukkan panjang element: ")
	fmt.Scan(&panjang)
	arr := make([]int, panjang)
	for i:=0 ; i< panjang ; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("\nNilai minimum dari array adalah ",min(arr...))
}
