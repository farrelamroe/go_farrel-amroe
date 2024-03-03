package main

import (
	"fmt"
)

// Struct untuk menyimpan data pelanggan
type Customer struct {
	Name          string
	Address       string
	RentedVehicles []Vehicle
}

// Struct untuk menyimpan data kendaraan
type Vehicle struct {
	Name  string
	Model string
	Price int
}

// Method untuk menampilkan data pelanggan
func (c Customer) Display() {
	fmt.Println("Nama Pelanggan:", c.Name)
	fmt.Println("Alamat:", c.Address)
}

// Method untuk menampilkan data kendaraan yang disewa oleh pelanggan
func (c Customer) DisplayRentedVehicles() {
	fmt.Println("Kendaraan yang disewa oleh", c.Name, ":")
	for _, v := range c.RentedVehicles {
		fmt.Printf("- %s %s\n", v.Name, v.Model)
	}
}

// Method untuk menyewa kendaraan
func (c *Customer) RentVehicle(v Vehicle) {
	c.RentedVehicles = append(c.RentedVehicles, v)
	fmt.Printf("%s %s berhasil disewa oleh %s\n", v.Name, v.Model, c.Name)
}

// Method untuk mengembalikan kendaraan yang disewa
func (c *Customer) ReturnVehicle(v Vehicle) {
	for i, rentedVehicle := range c.RentedVehicles {
		if rentedVehicle.Name == v.Name && rentedVehicle.Model == v.Model {
			c.RentedVehicles = append(c.RentedVehicles[:i], c.RentedVehicles[i+1:]...)
			fmt.Printf("%s %s berhasil dikembalikan oleh %s\n", v.Name, v.Model, c.Name)
			return
		}
	}
	fmt.Printf("%s %s tidak ditemukan dalam daftar kendaraan yang disewa oleh %s\n", v.Name, v.Model, c.Name)
}

func main() {
	// Meminta input data pelanggan
	var name, address string
	fmt.Print("Masukkan nama pelanggan: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan alamat pelanggan: ")
	fmt.Scanln(&address)

	// Membuat data pelanggan
	customer := Customer{
		Name:    name,
		Address: address,
	}

	// Menampilkan data pelanggan
	fmt.Println("----------------------")
	customer.Display()
	fmt.Println("----------------------")

	// Meminta input data kendaraan
	var vehicleName, vehicleModel string
	var vehiclePrice int
	fmt.Print("Masukkan nama kendaraan: ")
	fmt.Scanln(&vehicleName)
	fmt.Print("Masukkan model kendaraan: ")
	fmt.Scanln(&vehicleModel)
	fmt.Print("Masukkan harga sewa kendaraan: ")
	fmt.Scanln(&vehiclePrice)

	// Membuat data kendaraan
	vehicle := Vehicle{
		Name:  vehicleName,
		Model: vehicleModel,
		Price: vehiclePrice,
	}

	// Menyewa kendaraan oleh pelanggan
	customer.RentVehicle(vehicle)

	// Menampilkan data pelanggan setelah menyewa kendaraan
	fmt.Println("----------------------")
	customer.Display()
	customer.DisplayRentedVehicles()
	fmt.Println("----------------------")

	// Meminta input untuk mengembalikan kendaraan
	fmt.Print("Masukkan nama kendaraan yang ingin dikembalikan: ")
	fmt.Scanln(&vehicleName)
	fmt.Print("Masukkan model kendaraan yang ingin dikembalikan: ")
	fmt.Scanln(&vehicleModel)

	// Membuat data kendaraan yang akan dikembalikan
	returnedVehicle := Vehicle{
		Name:  vehicleName,
		Model: vehicleModel,
	}

	// Mengembalikan kendaraan oleh pelanggan
	customer.ReturnVehicle(returnedVehicle)

	// Menampilkan data pelanggan setelah mengembalikan kendaraan
	fmt.Println("----------------------")
	customer.Display()
	customer.DisplayRentedVehicles()
	fmt.Println("----------------------")
}
