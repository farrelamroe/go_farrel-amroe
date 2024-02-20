package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(str1, str2 string) bool {

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	runes1 := []rune(str1)
	runes2 := []rune(str2)

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	for i := range runes1 {
		if runes1[i] != runes2[i] {
			return false
		}
	}

	return true
}

func main() {
	var str1, str2 string
	fmt.Printf("Masukkan string 1: ")
	fmt.Scanf("%s", &str1)
	fmt.Printf("Masukkan string 2: ")
	fmt.Scanf("%s", &str2)

	if isAnagram(str1, str2) {
		fmt.Printf("%s dan %s adalah anagram.\n", str1, str2)
	} else {
		fmt.Printf("%s dan %s bukan anagram.\n", str1, str2)
	}
}
