package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

    var namaBuah string
    fmt.Print("Masukkan nama buah: ")
    fmt.Scanln(&namaBuah)

    url := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%s", namaBuah)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Data Nutrisi", namaBuah)
    fmt.Println("===================")
    fmt.Println("Kalori:", string(body))

    resp.Body.Close()
}