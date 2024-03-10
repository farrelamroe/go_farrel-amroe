package main

import "fmt"

func Romawi(number int) string {
	romanNumerals := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	divisors := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	res := ""
	for _, idx := range divisors {
		count := number / idx

		for i := 0; i < count; i++ {
			res += romanNumerals[idx]
		}

		number %= idx
	}

	return res
}

func main() {

	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	roman := Romawi(number)

	// Menampilkan hasil konversi
	fmt.Println("Angka Romawi:", roman)
}
