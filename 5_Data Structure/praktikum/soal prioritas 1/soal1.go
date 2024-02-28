package main

import (
	"fmt"
	"sort"
)

func flattenArray(arr [][]int) []int {
	var flatList []int
	uniqueItems := make(map[int]bool)

	for _, sublist := range arr {
		for _, item := range sublist {
			if !uniqueItems[item] {
				flatList = append(flatList, item)
				uniqueItems[item] = true
			}
		}
	}

	return flatList
}

func main() {
	inputArray := [][]int{{}, {}}

	for i := 0; i < 2; i++ {
		var sublistLength int
		fmt.Printf("Masukkan panjang sublist %d: ", i+1)
		fmt.Scanln(&sublistLength)

		inputArray[i] = make([]int, sublistLength)
		fmt.Printf("Masukkan elemen sublist %d: ", i+1)
		for j := 0; j < sublistLength; j++ {
			fmt.Scan(&inputArray[i][j])
		}
	}

	outputArray := flattenArray(inputArray)

	// Sort the output array
	sort.Ints(outputArray)

	fmt.Println("Output:", outputArray)
}
