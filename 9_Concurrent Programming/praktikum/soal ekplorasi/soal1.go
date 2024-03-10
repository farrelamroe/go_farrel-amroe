package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Struct untuk menampung data pengguna (users)
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	// Buat channel untuk menyimpan data pengguna (users)
	usersCh := make(chan []User)

	// Buat WaitGroup untuk sinkronisasi goroutine
	var wg sync.WaitGroup

	// Jalankan goroutine untuk mengambil data pengguna dari API
	wg.Add(1)
	go func() {
		defer wg.Done()
		getUsers(usersCh)
	}()

	// Terima dan cetak data pengguna dari channel
	users := <-usersCh
	for _, user := range users {
		fmt.Printf("ID: %d\nName: %s\nUsername: %s\nEmail: %s\n\n", user.ID, user.Name, user.Username, user.Email)
	}

	// Tunggu semua goroutine selesai
	wg.Wait()
}

// Fungsi untuk mengambil data pengguna dari API dan mengirimkannya ke dalam channel
func getUsers(usersCh chan<- []User) {
    // Kirim HTTP GET request ke API endpoint
    response, err := http.Get("https://fakestoreapi.com/users")
    if err != nil {
        fmt.Println("Error:", err)
        close(usersCh)
        return
    }
    defer response.Body.Close()

    // Baca response body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error:", err)
        close(usersCh)
        return
    }

    // Unmarshal data JSON ke dalam variabel kosong
    var rawData []map[string]interface{}
    err = json.Unmarshal(body, &rawData)
    if err != nil {
        fmt.Println("Error:", err)
        close(usersCh)
        return
    }

    // Konversi variabel kosong ke dalam struktur yang diinginkan
    var users []User
    for _, data := range rawData {
        user := User{
            ID:       int(data["id"].(float64)),
            Name:     data["name"].(string),
            Username: data["username"].(string),
            Email:    data["email"].(string),
        }
        users = append(users, user)
    }

    // Kirim data pengguna ke dalam channel
    usersCh <- users
}


