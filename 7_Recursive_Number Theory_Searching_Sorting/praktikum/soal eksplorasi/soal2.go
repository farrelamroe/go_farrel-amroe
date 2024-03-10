package main

import (
	"fmt"
	"sort"
)

func getItem(items []int, target int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(items)))

	collectedItems := []int{}
	sum := 0
	for _, item := range items {

		if sum+item <= target {
			collectedItems = append(collectedItems, item)
			sum += item
		}

		if sum == target {
			break
		}
	}

	return collectedItems
}

func main() {
	var target int
	fmt.Print("Masukkan target poin: ")
	fmt.Scanln(&target)

	var n int
	fmt.Print("Masukkan jumlah item: ")
	fmt.Scanln(&n)

	items := make([]int, n)
	fmt.Println("Masukkan poin setiap item:")
	for i := 0; i < n; i++ {
		fmt.Printf("Item %d: ", i+1)
		fmt.Scanln(&items[i])
	}

	fmt.Println("Item yang terkumpul:", getItem(items, target))
}
