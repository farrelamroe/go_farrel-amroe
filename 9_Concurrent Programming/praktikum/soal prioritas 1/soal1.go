package main

import (
	"fmt"
	"time"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func displayProcess(kata string) {
	for i := 1; i <= len(kata); i++ {
		fmt.Println(kata[:i])
		time.Sleep(3 * time.Second)
	}
}

func main() {
	kata := "kasur"
	reversedWord := reverseString(kata)
	go displayProcess(reversedWord)
	time.Sleep(time.Duration(len(reversedWord)) * 3 * time.Second)
}
