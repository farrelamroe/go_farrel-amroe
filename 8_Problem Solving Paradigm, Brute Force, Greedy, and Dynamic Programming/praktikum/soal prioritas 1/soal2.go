package main

import "fmt"

func generate(numRows int) [][]int {
    if numRows <= 0 {
        return [][]int{}
    }

    triangle := make([][]int, numRows)

    for i := 0; i < numRows; i++ {
        triangle[i] = make([]int, i+1)
        triangle[i][0] = 1
        triangle[i][i] = 1

        for j := 1; j < i; j++ {
            triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
    }

    return triangle
}

func main() {
    var numRows int
	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&numRows)
    pascalTriangle := generate(numRows)

    for _, row := range pascalTriangle {
        fmt.Println(row)
    }
}
