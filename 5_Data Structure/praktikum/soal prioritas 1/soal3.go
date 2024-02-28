package main

import (
	"fmt"
	"sort"
)

func sum(numbers []float64) float64 {
	sum := 0.0
	for _,num := range numbers {
		sum += num
	}
	return sum
}

func mean(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func median(numbers []float64) float64 {
	sort.Float64s(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2.0
	}
	return numbers[mid]
}

func mode(numbers []float64) float64 {
	freq := make(map[float64]int)
	for _, num := range numbers {
		freq[num]++
	}
	maxFreq := 0
	modeValue := 0.0
	for num, count := range freq {
		if count > maxFreq {
			maxFreq = count
			modeValue = num
		}
	}
	return modeValue
}

func main() {
	var n int
	fmt.Printf("Masukkan panjang list: ")
	fmt.Scan(&n)
	numbers := make([]float64, n)
	for num := range n {
		fmt.Scan(&numbers[num])
	}
	fmt.Println("Data:", numbers)
	fmt.Println("Sum: %2f\n", sum(numbers))
	fmt.Printf("Mean: %.2f\n", mean(numbers))
	fmt.Printf("Median: %.2f\n", median(numbers))
	fmt.Printf("Mode: %.2f\n", mode(numbers))
}
