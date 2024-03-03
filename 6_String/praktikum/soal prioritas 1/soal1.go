package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat: ")
	kalimat, _ := reader.ReadString('\n')

	fmt.Print("Huruf vokal dalam kalimat: ")
	vokal := "aeiouAEIOU"
	for _, char := range kalimat {
		if strings.ContainsRune(vokal, char) {
			fmt.Printf("%c", char)
		}
	}
}
