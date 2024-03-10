package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Fungsi untuk menghitung frekuensi kemunculan kata dari sebuah dokumen teks
func countWordFrequency(text string, wordFreq map[string]int, wg *sync.WaitGroup, ch chan<- map[string]int) {
	defer wg.Done()

	// Pisahkan teks menjadi kata-kata
	words := strings.Fields(text)

	// Hitung frekuensi kemunculan setiap kata
	for _, word := range words {
		wordFreq[word]++
	}

	ch <- wordFreq
}

func main() {
	// Buka file dokumen teks
	file, err := os.Open("9_Concurrent Programming/praktikum/soal prioritas 2/dokumen.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Baca isi file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	text := string(buffer)

	// Inisialisasi map untuk menyimpan frekuensi kemunculan kata
	wordFreq := make(map[string]int)

	// Buat channel untuk menyimpan hasil perhitungan frekuensi dari setiap goroutine
	ch := make(chan map[string]int)

	// Buat WaitGroup untuk sinkronisasi goroutine
	var wg sync.WaitGroup

	// Pisahkan teks menjadi beberapa bagian
	parts := strings.Split(text, "\n")

	// Proses setiap bagian teks secara concurrent
	for _, part := range parts {
		wg.Add(1)
		go countWordFrequency(part, make(map[string]int), &wg, ch)
	}

	// Goroutine untuk mengumpulkan hasil perhitungan frekuensi dari channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Terima hasil perhitungan frekuensi dan tambahkan ke dalam map wordFreq
	for result := range ch {
		for word, freq := range result {
			wordFreq[word] += freq
		}
	}

	// Tulis hasil perhitungan frekuensi ke dalam file CSV
	csvFile, err := os.Create("word_frequency.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Tulis header
	writer.Write([]string{"word", "frequency"})

	// Tulis setiap kata dan frekuensinya ke dalam file CSV
	for word, freq := range wordFreq {
		writer.Write([]string{word, fmt.Sprintf("%d", freq)})
	}

	fmt.Println("Perhitungan frekuensi kata telah disimpan dalam file word_frequency.csv")
}
