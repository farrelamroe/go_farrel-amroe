package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func groupPalindromes(words []string) [][]string {
	groups := make(map[int][]string)
	for _, word := range words {
		length := len(word)
		if isPalindrome(word) {
			groups[length] = append(groups[length], word)
		}
	}

	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	var input string
	fmt.Print("Masukkan kata-kata (pisahkan dengan koma): ")
	fmt.Scanln(&input)

	words := strings.Split(input, ",")

	fmt.Println("Kelompok kata palindrom:")
	grouped := groupPalindromes(words)
	for _, group := range grouped {
		fmt.Println(group)
	}
}
