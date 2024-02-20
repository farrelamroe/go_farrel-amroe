package main

import "fmt"

func main() {
	var nilai int
	fmt.Printf("Masukkan nilai: " )
	fmt.Scanf("%d", &nilai)
	if nilai >= 85 && nilai <= 100 {
		fmt.Println("A")
	} else if nilai >= 70 && nilai <= 84 {
		fmt.Println("B")
	} else if nilai >= 55 && nilai <= 69 {
		fmt.Println("C")
	} else if nilai >= 40 && nilai <= 44 {
		fmt.Println("D")
	} else if nilai >= 0 && nilai <= 39 {
		fmt.Println("E")
	} else {
		fmt.Println("Nilai invalid")
	}
}
