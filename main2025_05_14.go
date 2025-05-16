// challege 1

// package main

// import "fmt"

// type Team struct {
// 	Name  string
// 	Score int
// }

// func newTeam(name string, score int) *Team {
// 	return &Team{
// 		Name:  name,
// 		Score: score,
// 	}
// }

// func updatedScore(team *Team, score int) {
// 	fmt.Printf("Data tim sebelum update: %+v\n", team)
// 	if team.Score < score {
// 		fmt.Printf("Memperbarui skor %s dari %d menjadi %d\n", team.Name, team.Score, score)
// 		team.Score = score
// 		fmt.Println("Skor berhasil diupdate!")
// 	} else {
// 		fmt.Printf("Skor %s tidak diperbarui karena skor baru (%d) tidak lebih besar dari skor saat ini (%d)\n", team.Name, score, team.Score)
// 		fmt.Println("Tidak ada update skor.")
// 	}
// }

// func main() {
// 	// Inisialisasi
// 	team := newTeam("Persib", 70)
// 	updatedScore(team, 80)
// 	fmt.Printf("Data tim setelah update: %+v\n", team)
// }

// challege 2
// package main

// import "fmt"

// type Buku struct {
// 	Judul      string
// 	Penulis    string
// 	TahunTerbit int
// }

// func filterBukuTahun(buku []Buku, tahun int) []Buku {
// 	var bukuFilter []Buku
// 	for _, b := range buku {
// 		if b.TahunTerbit == tahun {
// 			bukuFilter = append(bukuFilter, b)
// 		}
// 	}
// 	return bukuFilter
// }

// func main() {

// 	daftarBuku := []Buku{
// 		{Judul: "Indonesia Raya", Penulis: "Apa", TahunTerbit: 1945},
// 		{Judul: "Jaya Indonesiaku", Penulis: "Api", TahunTerbit: 1945},
// 		{Judul: "Indonesia Jaya", Penulis: "Apo", TahunTerbit: 1954},
// 		{Judul: "Indo Asia", Penulis: "Apu", TahunTerbit: 1955},
// 		{Judul: "Asia asia", Penulis: "Ape", TahunTerbit: 1994},
// 	}

// 	tahunFilter := 1945

// 	bukuSesuaiTahun := filterBukuTahun(daftarBuku, tahunFilter)

// 	fmt.Printf("Daftar buku yang terbit pada tahun %d:\n", tahunFilter)
// 	for _, b := range bukuSesuaiTahun {
// 		fmt.Printf("- %s (%s, %d)\n", b.Judul, b.Penulis, b.TahunTerbit)
// 	}
// }

// challenge 3
// package main

// import "fmt"

// type PaymentMethod interface {
// 	TutorialPembayaran() string
// }

// // Qris
// type QRIS struct {
// 	Nama string
// }

// // EWallet
// type EWallet struct {
// 	Nama string
// }

// // VirtualAccount
// type VirtualAccount struct {
// 	Bank string
// }

// // TutorialPembayaran Qris
// func (q QRIS) TutorialPembayaran() string {
// 	return `
// 	Tutorial Pembayaran Qris:
// 	1. Buka aplikasi Pembayaran Qris.
// 	2. Pilih menu "Bayar" atau "Scan".
// 	3. Scan QR Code.
// 	4. Masukkan Nominal.
// 	5. Masukkan Pinconst.
// 	`
// }

package main

import "fmt"

// PaymentMethod adalah interface untuk metode pembayaran
type PaymentMethod interface {
	TutorialPembayaran() string
}

// QRIS adalah struct untuk metode pembayaran QRIS
type QRIS struct {
	Nama string
}

// EWallet adalah struct untuk metode pembayaran E-Wallet
type EWallet struct {
	Nama string
}

// VirtualAccount adalah struct untuk metode pembayaran Virtual Account
type VirtualAccount struct {
	Bank string
}

// TutorialPembayaran untuk QRIS
func (q QRIS) TutorialPembayaran() string {
	return `
	Tutorial Pembayaran QRIS:
	1. Buka aplikasi pembayaran (misalnya, GoPay, OVO, Dana, atau aplikasi bank).
	2. Pilih menu "Bayar" atau "Scan".
	3. Scan QR Code yang tersedia.
	4. Masukkan nominal pembayaran.
	5. Konfirmasi pembayaran dan masukkan PIN.
	6. Pembayaran berhasil!
	`
}

// TutorialPembayaran untuk E-Wallet
func (e EWallet) TutorialPembayaran() string {
	return `
	Tutorial Pembayaran E-Wallet:
	1. Buka aplikasi E-Wallet (misalnya, GoPay, OVO, Dana).
	2. Pilih menu "Bayar".
	3. Masukkan nomor telepon yang terdaftar pada E-Wallet.
	4. Masukkan nominal pembayaran.
	5. Konfirmasi pembayaran dan masukkan PIN.
	6. Pembayaran berhasil!
	`
}

// TutorialPembayaran untuk Virtual Account
func (v VirtualAccount) TutorialPembayaran() string {
	return fmt.Sprintf(`
	Tutorial Pembayaran Virtual Account (%s):
	1. Buka aplikasi bank Anda atau ATM.
	2. Pilih menu "Transfer".
	3. Pilih "Transfer ke Virtual Account" atau "Pembayaran Virtual Account".
	4. Masukkan nomor Virtual Account yang tertera.
	5. Masukkan nominal pembayaran.
	6. Konfirmasi pembayaran.
	7. Pembayaran berhasil!
	`, v.Bank)
}

// TampilkanTutorial menampilkan tutorial pembayaran berdasarkan metode yang dipilih
func TampilkanTutorial(metode PaymentMethod) {
	fmt.Println(metode.TutorialPembayaran())
}

func main2025_05_14() {
	// Contoh penggunaan
	qris := QRIS{Nama: "QRIS"}
	ewallet := EWallet{Nama: "OVO"}
	virtualAccount := VirtualAccount{Bank: "BCA"}

	fmt.Println("Pilih metode pembayaran:")
	fmt.Println("1. QRIS")
	fmt.Println("2. E-Wallet")
	fmt.Println("3. Virtual Account")

	var pilihan int
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		TampilkanTutorial(qris)
	case 2:
		TampilkanTutorial(ewallet)
	case 3:
		TampilkanTutorial(virtualAccount)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
