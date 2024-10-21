package main

import (
	"fmt"
	"strings"
)

// Struct untuk menyimpan informasi tentang destinasi wisata
type Destinasi struct {
	Nama      string
	Lokasi    string
	Deskripsi string
	Biaya     float64
	Fasilitas string
}

// Struct untuk menyimpan informasi pengguna
type User struct {
	Username string
	Password string
}

const maxDestinasi = 100
const maxUsers = 100

// variabel
var isAdmin bool
var destinasiList [maxDestinasi]Destinasi
var destinasiCount int
var userList [maxUsers]User
var userCount int

func main() {
	initUsers() // Inisialisasi pengguna default
	login()
	for {
		fmt.Println("\nSelamat datang di Aplikasi Pariwisata")
		fmt.Println("=======================================")
		if isAdmin {
			fmt.Println("1. Tambah Tempat Wisata")
			fmt.Println("2. Ubah Tempat Wisata")
			fmt.Println("3. Hapus Tempat Wisata")
			fmt.Println("4. Tampilkan Daftar Tempat Wisata")
			fmt.Println("5. Keluar")
		} else {
			fmt.Println("1. Tampilkan Daftar Tempat Wisata")
			fmt.Println("2. Cari Tempat Wisata")
			fmt.Println("3. Ganti User")
			fmt.Println("4. Keluar")
		}

		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			if isAdmin {
				addTempatWisata()
			} else {
				showTempatWisata()
			}
		case 2:
			if isAdmin {
				editTempatWisata()
			} else {
				searchTempatWisata()
			}
		case 3:
			if isAdmin {
				deleteTempatWisata()
			} else {
				changeUser()
			}
		case 4:
			if isAdmin {
				showTempatWisata()
			} else {
				fmt.Println("Terima kasih telah menggunakan aplikasi!")
				return
			}
		case 5:
			if isAdmin {
				login()
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func initUsers() {
	// Tambahkan admin default
	userList[0] = User{Username: "admin", Password: "admin123"}
	userCount = 1
}

func login() {
	for {
		fmt.Println("\nLogin")
		fmt.Println("1. Login")
		fmt.Println("2. Daftar Pengguna Baru")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var username, password string
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)

			for i := 0; i < userCount; i++ {
				if userList[i].Username == username && userList[i].Password == password {
					isAdmin = (username == "admin")
					if isAdmin {
						fmt.Println("Login sebagai Admin berhasil")
					} else {
						fmt.Println("Login sebagai Pengguna berhasil")
					}
					return
				}
			}

			fmt.Println("Login gagal. Username atau password salah")
		case 2:
			addUser()
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Fungsi untuk menambahkan tempat wisata baru
func addTempatWisata() {
	if destinasiCount >= maxDestinasi {
		fmt.Println("Tidak dapat menambahkan tempat wisata lagi, kapasitas penuh.")
		return
	}

	var nama, lokasi, deskripsi, fasilitas string
	var biaya float64

	fmt.Print("Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Lokasi: ")
	fmt.Scanln(&lokasi)

	fmt.Print("Deskripsi: ")
	fmt.Scanln(&deskripsi)

	fmt.Print("Biaya: ")
	fmt.Scanln(&biaya)

	fmt.Print("Fasilitas: ")
	fmt.Scanln(&fasilitas)

	newDestinasi := Destinasi{Nama: nama, Lokasi: lokasi, Deskripsi: deskripsi, Biaya: biaya, Fasilitas: fasilitas}
	destinasiList[destinasiCount] = newDestinasi
	destinasiCount++

	fmt.Println("Tempat wisata baru telah ditambahkan!")

	// Urutkan destinasiList setelah menambahkan elemen baru
	insertionSort(destinasiList[:destinasiCount])
}

// Fungsi untuk mengubah data tempat wisata
func editTempatWisata() {
	var namaLama, namaBaru, lokasiBaru, deskripsiBaru, fasilitasBaru string
	var biayaBaru float64

	fmt.Print("Nama tempat wisata yang ingin diubah: ")
	fmt.Scanln(&namaLama)

	index := binarySearch(namaLama)
	if index == -1 {
		fmt.Println("Tempat wisata tidak ditemukan!")
		return
	}

	fmt.Print("Nama baru: ")
	fmt.Scanln(&namaBaru)

	fmt.Print("Lokasi baru: ")
	fmt.Scanln(&lokasiBaru)

	fmt.Print("Deskripsi baru: ")
	fmt.Scanln(&deskripsiBaru)

	fmt.Print("Biaya baru: ")
	fmt.Scanln(&biayaBaru)

	fmt.Print("Fasilitas baru: ")
	fmt.Scanln(&fasilitasBaru)

	destinasiList[index].Nama = namaBaru
	destinasiList[index].Lokasi = lokasiBaru
	destinasiList[index].Deskripsi = deskripsiBaru
	destinasiList[index].Biaya = biayaBaru
	destinasiList[index].Fasilitas = fasilitasBaru

	fmt.Println("Data tempat wisata telah diubah!")

	// Urutkan ulang destinasiList setelah mengubah elemen
	insertionSort(destinasiList[:destinasiCount])
}

// Fungsi untuk menghapus tempat wisata
func deleteTempatWisata() {
	var nama string
	fmt.Print("Nama tempat wisata yang ingin dihapus: ")
	fmt.Scanln(&nama)

	index := binarySearch(nama)
	if index == -1 {
		fmt.Println("Tempat wisata tidak ditemukan!")
		return
	}

	copy(destinasiList[index:], destinasiList[index+1:destinasiCount])
	destinasiCount--
	destinasiList[destinasiCount] = Destinasi{} // Hapus elemen terakhir

	fmt.Println("Tempat wisata telah dihapus!")
}

// Fungsi untuk menampilkan daftar tempat wisata
func showTempatWisata() {
	if destinasiCount == 0 {
		fmt.Println("Belum ada tempat wisata yang ditambahkan.")
		return
	}
	fmt.Println("\nDaftar Tempat Wisata:")
	for i := 0; i < destinasiCount; i++ {
		dest := destinasiList[i]
		fmt.Printf("%d. %s (%s)\n", i+1, dest.Nama, dest.Lokasi)
		fmt.Printf("   Deskripsi: %s\n", dest.Deskripsi)
		fmt.Printf("   Biaya: %.2f\n", dest.Biaya)
		fmt.Printf("   Fasilitas: %s\n", dest.Fasilitas)
	}
}

// Fungsi untuk mencari tempat wisata berdasarkan kata kunci
func searchTempatWisata() {
	var keyword string
	fmt.Print("Masukkan kata kunci pencarian(Nama/Lokasi/Deskripsi/Fasilitas): ")
	fmt.Scanln(&keyword)

	keyword = strings.ToLower(keyword)
	found := false

	fmt.Println("\nHasil Pencarian:")
	for i := 0; i < destinasiCount; i++ {
		dest := destinasiList[i]
		if strings.Contains(strings.ToLower(dest.Nama), keyword) ||
			strings.Contains(strings.ToLower(dest.Lokasi), keyword) ||
			strings.Contains(strings.ToLower(dest.Deskripsi), keyword) ||
			strings.Contains(strings.ToLower(dest.Fasilitas), keyword) {
			fmt.Printf("%d. %s (%s)\n", i+1, dest.Nama, dest.Lokasi)
			fmt.Printf("   Deskripsi: %s\n", dest.Deskripsi)
			fmt.Printf("   Biaya: %.2f\n", dest.Biaya)
			fmt.Printf("   Fasilitas: %s\n", dest.Fasilitas)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada tempat wisata yang sesuai dengan kata kunci.")
	}
}

// Fungsi untuk mengganti user
func changeUser() {
	login() // Memanggil fungsi login untuk meminta username dan password lagi
}

// Fungsi untuk menambahkan pengguna baru
func addUser() {
	if userCount >= maxUsers {
		fmt.Println("Tidak dapat menambahkan pengguna lagi, kapasitas penuh.")
		return
	}

	var username, password string
	fmt.Print("Masukkan username baru: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password baru: ")
	fmt.Scanln(&password)

	for i := 0; i < userCount; i++ {
		if userList[i].Username == username {
			fmt.Println("Username sudah ada, silakan pilih username lain.")
			return
		}
	}

	newUser := User{Username: username, Password: password}
	userList[userCount] = newUser
	userCount++
	fmt.Println("Pengguna baru telah ditambahkan!")
}

// Fungsi untuk mengurutkan array destinasiList menggunakan Insertion Sort
func insertionSort(destinations []Destinasi) {
	for i := 1; i < len(destinations); i++ {
		key := destinations[i]
		j := i - 1
		for j >= 0 && strings.ToLower(destinations[j].Nama) > strings.ToLower(key.Nama) {
			destinations[j+1] = destinations[j]
			j--
		}
		destinations[j+1] = key
	}
}

// Fungsi untuk melakukan binary search pada array destinasiList
func binarySearch(nama string) int {
	low := 0
	high := destinasiCount - 1
	nama = strings.ToLower(nama)

	for low <= high {
		mid := (low + high) / 2
		midNama := strings.ToLower(destinasiList[mid].Nama)
		if midNama == nama {
			return mid
		} else if midNama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
