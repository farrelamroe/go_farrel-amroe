package main

import "fmt"

func triangle(n int) int {
	return n*(n+1)/2
}

func main(){
	var number int
	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&number)
	fmt.Println(triangle(number))
}