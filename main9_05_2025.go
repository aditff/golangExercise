package main

import (
	"fmt"
)

// func main() {
// 	// Deklarasi array dengan 4 elemen bertipe int, semua default (0)
// 	arr := [4]int{}
// 	arr[3] = 2 // Mengisi elemen ke-4 (indeks 3) dengan nilai 2
// 	fmt.Println("Array:", arr)

// 	// Deklarasi slice dengan 6 elemen
// 	slice := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println("Data pertama:", slice)

// 	// Menambahkan elemen ke slice menggunakan append
// 	slice = append(slice, 7)
// 	fmt.Println("Data kedua:", slice)

// 	// Menyalin 4 elemen pertama dari slice ke slice2
// 	slice2 := make([]int, 4)
// 	copy(slice2, slice)
// 	fmt.Println("Data slice2:", slice2)

// 	// Menampilkan panjang dan kapasitas slice
// 	fmt.Println("Len:", len(slice))
// 	fmt.Println("Cap:", cap(slice))
// }

// func main() {
// 	// for classic
// 	// for i := 0; i <= 5; i++ {
// 	// 	fmt.Println("lumoshive academy index :", i)
// 	// }

// 	// for infinite
// 	// for {
// 	// 	fmt.Println("hhalloo")
// 	// }

// 	// for range
// 	// slice := []string{"Hallo", "lumoshive", "academy"}
// 	// for index, str := range slice {
// 	// 	fmt.Println("data :", str, "index :", index)
// 	// }

// 	// for statement
// 	// counter := 0
// 	// for counter < 5 {
// 	// 	fmt.Println("Counter :", counter)
// 	// 	counter++
// 	// }
// }

// // basic
// // func geet() {
// // 	fmt.Println("hello Lumoshive")
// // }

// condition
// func checkGrade(num int) string {
// 	if num > 90 {
// 		return "lulus"
// 	}

// 	return "tidak lulus"
// }

// if

// func main() {
// 	// grade := checkGrade(6)
// 	// println("grade :", grade)

// 	grade := checkGrade(6, "")
// 	println("grade :", grade)

// 	day := checkDay(3)
// 	println(day)

// }

// func checkGrade(num int, status string) string {
// 		if num > 90 && status == "aktif" {
// 			return "A"
// 		}

// 		return "B"
// 	}

// 	// swicth
// 	func checkDay(num int) string {
// 		switch num {
// 		case 1:
// 			return "senin"
// 					case 2:
// 			return "selasa"
// 					case 3:
// 			return "rabu"
// 		default:
// 			return "hari tidak valid"
// 		}
// 	}

// var num int = 1

// var numCopy int
// func main() {
// 	// numCopy := num
// 	// println(numCopy)

// 	// var pointer *int
// 	// pointer = &num
// 	// println(pointer)

// 	// var pointer *int = &num
// 	// println(pointer)
// 	// println(*pointer)
// 	// println(&num)

// 	numCopy = num
// 	numCopy = 7

// 	println(numCopy)
// 	println(num)

// 	var pointer *int = &num
// 	println(num)
// 	*pointer = 5
// 	println(num)
// }

// func main() {

// 	data :=[]int{5, 6, 7, 9 , 15, 10 , 3 , 5 , 6, 7, 8, 4, 8, 2, 1, 4, 9, 15, 27, 10, 1}

// 	// panjang slice 18
// 	startIndex := 6
// 	endIndex := 18

// 	dataReturn:= data[startIndex: endIndex]
// 	fmt.Println(dataReturn)
// }

// func main() {
// 	// Data slice
// 	data := []int{5, 6, 7, 9, 15, 10, 3, 5, 6, 7, 8, 2, 1, 4, 9, 15, 27, 10, 1}

// 	// Indeks yg di hitung rata-ratanya
// 	indexes := []int{5, 6, 8, 10, 15}

// 	// Hitung rata-rata
// 	total := 0
// 	for _, idx := range indexes {
// 		// Cek agar index tidak out-of-range
// 		if idx >= 0 && idx < len(data) {
// 			total += data[idx]
// 		} else {
// 			fmt.Printf("Index %d di luar batas slice\n", idx)
// 			return
// 		}
// 	}

// 	average := total / len(indexes)

// 	fmt.Println("Nilai rata-rata dari index 5, 6, 8, 10, 15 adalah", average)
// }

// Function untuk menghitung total pembayaran setelah diskon
func hitungPembayaran(jumlahBarang int) float64 {
    hargaSatuan := 100000.0 // Harga satuan Rp 100.000
    totalHarga := float64(jumlahBarang) * hargaSatuan
    var diskon float64

    // Menentukan diskon berdasarkan jumlah barang
    switch {
    case jumlahBarang == 2:
        diskon = 0.10 // 10%
    case jumlahBarang == 4:
        diskon = 0.50 // 50%
    case jumlahBarang > 4:
        diskon = 0.75 // 75%
    default:
        diskon = 0.0 // Tidak ada diskon
    }

    // Menghitung total pembayaran setelah diskon
    totalBayar := totalHarga * (1 - diskon)
    return totalBayar
}

func main9_05_2025() {
    // Contoh penggunaan function
    var jumlah int

    fmt.Print("Masukkan jumlah barang yang dibeli: ")
    fmt.Scan(&jumlah)

    total := hitungPembayaran(jumlah)
    fmt.Printf("Jumlah pembayaran untuk %d barang adalah Rp %.0f\n", jumlah, total)
}