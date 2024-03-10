package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
    // Mencari biaya minimum untuk mencapai setiap batu
    n := len(cost)
    minCost := make([]int, n)
    minCost[0] = cost[0]
    minCost[1] = cost[1]

    for i := 2; i < n; i++ {
        minCost[i] = cost[i] + min(minCost[i-1], minCost[i-2])
    }

    // Mengembalikan biaya minimum untuk mencapai Batu N
    return min(minCost[n-1], minCost[n-2])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    // Meminta input dari pengguna
    var n int
    fmt.Print("Masukkan jumlah batu: ")
    fmt.Scan(&n)

    cost := make([]int, n)
    fmt.Println("Masukkan biaya untuk setiap batu: ")
    for i := 0; i < n; i++ {
        fmt.Scan(&cost[i])
    }

    // Memanggil fungsi minCostClimbingStairs dengan input dari pengguna
    fmt.Println("Biaya minimum yang diperlukan:", minCostClimbingStairs(cost))
}
