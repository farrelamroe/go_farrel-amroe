package main

import (
    "fmt"
)

type Stack struct {
    items map[string]bool
}

func NewStack() *Stack {
    return &Stack{items: make(map[string]bool)}
}

func (s *Stack) Push(item string) {
    if _, exists := s.items[item]; !exists {
        s.items[item] = true
    } else {
        fmt.Println("Data sudah ada dalam stack, tidak bisa menambahkan data yang sama")
    }
}

func (s *Stack) Pop() string {
    var item string
    for item = range s.items {
        delete(s.items, item)
        break
    }
    return item
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack) Values() []string {
    values := make([]string, len(s.items))
    i := 0
    for key := range s.items {
        values[i] = key
        i++
    }
    return values
}

func main() {
    stack := NewStack()

    for {
        var input string
        fmt.Print("Masukkan data (ketik 'selesai' untuk berhenti): ")
        fmt.Scanln(&input)

        if input == "selesai" {
            break
        }

        stack.Push(input)
    }

    fmt.Println("Nilai dalam stack:")
    for _, value := range stack.Values() {
        fmt.Println(value)
    }

    fmt.Println("Mengeluarkan nilai dari stack:")
    for !stack.IsEmpty() {
        fmt.Println(stack.Pop())
    }
}
