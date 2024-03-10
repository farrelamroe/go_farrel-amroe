package main

import "fmt"

func factorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n*factorial(n-1)
	}


}

func catalan(n int) int {
	return factorial(2*n)/(factorial(n+1)*factorial(n))
}

func main(){
	var number int
	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&number)
	fmt.Println(catalan(number))
}