package main

import "fmt"

func main() {
	var budget, waktu, kesulitan int
	skor := 0
	fmt.Printf("Masukkan budget: ")
	fmt.Scanf("%d", &budget)
	fmt.Printf("Masukkan waktu pengerjaan: ")
	fmt.Scanf("%d", &waktu)
	fmt.Printf("Masukkan tingkat kesulitan: ")
	fmt.Scanf("%d", &kesulitan)
	//Perhitungan budget
	if budget >= 0 && budget <= 50 {
		skor += 4
	} else if budget >= 51 && budget <= 80 {
		skor += 3
	} else if budget >= 81 && budget <= 100 {
		skor += 2
	} else if budget > 100 {
		skor += 1
	}
	//Perhitungan waktu pengerjaan
	if waktu >= 0 && waktu <= 20 {
		skor += 10
	} else if waktu >= 21 && waktu <= 30 {
		skor += 5
	} else if waktu >= 31 && waktu <= 50 {
		skor += 2
	} else if waktu > 50 {
		skor += 1
	}
	//Perhitungan tingkat kesulitan
	if kesulitan >= 0 && kesulitan <= 3 {
		skor += 10
	} else if kesulitan >= 4 && kesulitan <= 6 {
		skor += 5
	} else if kesulitan >= 7 && kesulitan <= 10 {
		skor += 1
	} else if kesulitan > 10 {
		skor += 0
	}
	//Perhitungan skor
	if skor >= 17 && skor <= 24 {
		fmt.Println("High")
	} else if skor >= 10 && skor <= 16 {
		fmt.Println("Medium")
	} else if skor >= 3 && skor <= 9 {
		fmt.Println("Low")
	} else if skor <=2 {
		fmt.Println("Impossible")
	}
}
