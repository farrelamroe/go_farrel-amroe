package main

import "fmt"

// Set adalah struktur data untuk merepresentasikan himpunan bilangan bulat unik
type Set struct {
	dt map[int]struct{} // Menggunakan map[int]struct{} karena tidak membutuhkan nilai untuk himpunan bilangan bulat unik
}

// Add digunakan untuk menambahkan bilangan bulat ke dalam himpunan
func (s *Set) Add(i int) {
	s.dt[i] = struct{}{}
}

// Get digunakan untuk mendapatkan seluruh bilangan bulat dalam himpunan
func (s Set) Get() []int {
	result := make([]int, 0, len(s.dt))

	for k := range s.dt {
		result = append(result, k)
	}

	return result
}

// Remove digunakan untuk menghapus bilangan bulat dari himpunan
func (s *Set) Remove(k int) {
	delete(s.dt, k)
}

func main() {
	s := Set{
		dt: make(map[int]struct{}),
	}

	s.Add(1)
	s.Add(2)
	s.Add(1) // Tidak akan ada efek karena 1 sudah ada di dalam himpunan
	s.Add(3)
	s.Add(4)
	s.Add(5)

	s.Remove(100) // Tidak akan ada efek karena 100 tidak ada di dalam himpunan

	fmt.Println(s.Get())
}
