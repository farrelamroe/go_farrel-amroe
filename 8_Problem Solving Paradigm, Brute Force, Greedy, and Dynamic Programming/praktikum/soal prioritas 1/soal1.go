package main

import (
	"fmt"
	"strconv"
)

func binaryRepresentation(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	var input int
	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&input)
	output := binaryRepresentation(input)
	fmt.Println("Output:", output)
}
