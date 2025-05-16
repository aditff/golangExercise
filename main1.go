package main

import (
	"fmt"
)

// var num int
// var name string
// var status bool
// var total float64

// func main() {
// 	println(num)
// 	println(name)
// 	println(status)
// 	println(total)
// }

// const (
// 	// Read = 1 << iota
// 	Read = iota + 1
// 	Write
// 	Delete
// 	Execute
// )

// func main() {
// 	println(Read)
// 	println(Write)
// 	println(Delete)
// 	println(Execute)
// }

// var num int

// func main() {
// 	num_2 := "Number Satu"
// 	println(num)
// 	println(num_2)
// }

// func setNum() {
// 	num = 2
// }

// type CustomString = string

// func main() {
// 	var name CustomString = "acaddemy"
// 	println(name)
// }

// type alias
// type CustomString = string

// func main() {
// 	var name CustomString = "acaddemy"
// 	var second_name string = name
// 	println(second_name)
// }

// type CustomString = string // alias
// type UserID int // declaration

// func main() {
// 	var name CustomString = "acaddemy"
// 	var second_name CustomString = name
// 	println(second_name)

// 	var id UserID = 435
// 	var user_id int = int(id)
// 	println(user_id)
// }

// latihan
// func main() {
// 	kalimat := "Aku seorang programmer golang yang handal"

// 	byteS := []byte(kalimat)

// 	fmt.Println("Kalimat asli:")
// 	fmt.Println(kalimat)

// 	fmt.Println("Hasil konversi ke byte:")
// 	fmt.Println(byteS)
// }

// Aritmatika
// func main() {
// 	var A int = 30
// 	var B float64 = 24.5
// 	var C int = -45
// 	var D float64 = 0.67

// 	// var E = (A + B) * C / D
// 	var E = (float64(A) + B) * float64(C) / D

// 	fmt.Println("Hasil E =", E)
// }




func main1() {
	jumlahProduk := 100
	hargaJual := 150000
	hargaBeli := 100000
	biayaOperasionalPerPcs := 1000
	diskonPersen := 15

	hargaSetelahDiskon := hargaJual * (100 - diskonPersen) / 100
	totalPendapatan := hargaSetelahDiskon * jumlahProduk
	totalBiaya := (hargaBeli + biayaOperasionalPerPcs) * jumlahProduk
	totalKeuntungan := totalPendapatan - totalBiaya

	fmt.Println("Harga jual setelah diskon: rp.", hargaSetelahDiskon)
	fmt.Println("Total pendapatan: rp.", totalPendapatan)
	fmt.Println("Total biaya: rp.", totalBiaya)
	fmt.Println("Total keuntungan: rp.", totalKeuntungan)
}