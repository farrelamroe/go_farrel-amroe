package main

import "fmt"

func diamond(n int) int {
	return n*n
}

func main(){
	var number int
	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&number)
	fmt.Println(diamond(number))
}