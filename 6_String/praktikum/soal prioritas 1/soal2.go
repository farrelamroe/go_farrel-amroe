package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	var s string
	fmt.Printf("Masukkan kata: ")
	fmt.Scan(&s)
	res := sort.Strings(s)
	fmt.Println(res)
}
