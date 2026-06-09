package main

import "fmt"

const NMAX int = 100

type Perangkat struct {
	nama    string
	ruangan string
	watt    float64
	durasi  float64 
}

type TabPerangkat struct {
	data [NMAX]Perangkat
	n    int
}

// ==========================================
// BAGIAN A & B: MANIPULASI DATA (CRUD)
// ==========================================

// Menambahkan data perangkat baru
func tambahPerangkat(T *TabPerangkat, nama string, ruangan string, watt float64, durasi float64) {
	if T.n < NMAX {
		T.data[T.n].nama = nama
		T.data[T.n].ruangan = ruangan
		T.data[T.n].watt = watt
		T.data[T.n].durasi = durasi
		T.n = T.n + 1
		fmt.Println("Data perangkat berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas penyimpanan penuh!")
	}
}

// Mengubah data perangkat berdasarkan indeks
func ubahPerangkat(T *TabPerangkat, index int, nama string, ruangan string, watt float64, durasi float64) {
	if index >= 0 && index < T.n {
		T.data[index].nama = nama
		T.data[index].ruangan = ruangan
		T.data[index].watt = watt
		T.data[index].durasi = durasi
		fmt.Println("Data perangkat berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

// Menghapus data perangkat dengan menggeser elemen array
func hapusPerangkat(T *TabPerangkat, index int) {
	var i int
	if index >= 0 && index < T.n {
		for i = index; i < T.n-1; i++ {
			T.data[i] = T.data[i+1]
		}
		T.n = T.n - 1
		fmt.Println("Data perangkat berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

// Prosedur pembantu untuk menampilkan seluruh isi array ke layar
func cetakSemuaPerangkat(T TabPerangkat) {
	var i int
	if T.n == 0 {
		fmt.Println("(Tidak ada data perangkat)")
		return
	}
	for i = 0; i < T.n; i++ {
		fmt.Printf("%d. [%s] %s - %.1f Watt, Durasi: %.1f jam/hari\n", i, T.data[i].ruangan, T.data[i].nama, T.data[i].watt, T.data[i].durasi)
	}
}

// ==========================================
// BAGIAN C: PENCARIAN DATA
// ==========================================

// Pencarian berdasarkan Jenis Ruangan menggunakan Sequential Search
func cariBerdasarkanRuangan(T TabPerangkat, targetRuangan string) {
	var i int
	var ketemu bool

	ketemu = false
	fmt.Printf("\n--- Hasil Pencarian di Ruangan: %s ---\n", targetRuangan)

	for i = 0; i < T.n; i++ {
		if T.data[i].ruangan == targetRuangan {
			fmt.Printf("- %s (%.2f Watt, %.2f Jam/hari)\n", T.data[i].nama, T.data[i].watt, T.data[i].durasi)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada perangkat yang ditemukan di ruangan tersebut.")
	}
}

// Pencarian berdasarkan Nama menggunakan Binary Search
func cariBerdasarkanNama(T TabPerangkat, targetNama string) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = T.n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if T.data[tengah].nama == targetNama {
			return tengah // Mengembalikan indeks jika ketemu
		} else if T.data[tengah].nama < targetNama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1 // Mengembalikan -1 jika tidak ditemukan
}

// ==========================================
// BAGIAN D: PENGURUTAN DATA
// ==========================================

// Mengurutkan berdasarkan Konsumsi Energi TERTINGGI (Descending) dengan Selection Sort
func urutEnergiTertinggi(T *TabPerangkat) {
	var i, j, maxIdx int
	var temp Perangkat
	var energi1, energi2 float64

	for i = 0; i < T.n-1; i++ {
		maxIdx = i
		for j = i + 1; j < T.n; j++ {
			energi1 = T.data[j].watt * T.data[j].durasi
			energi2 = T.data[maxIdx].watt * T.data[maxIdx].durasi

			if energi1 > energi2 {
				maxIdx = j
			}
		}
		// Proses pertukaran data (Swap)
		temp = T.data[i]
		T.data[i] = T.data[maxIdx]
		T.data[maxIdx] = temp
	}
}

// Mengurutkan berdasarkan Abjad Nama (Ascending) dengan Insertion Sort
func urutAbjadNama(T *TabPerangkat) {
	var i, j int
	var temp Perangkat

	for i = 1; i < T.n; i++ {
		temp = T.data[i]
		j = i - 1

		for j >= 0 && T.data[j].nama > temp.nama {
			T.data[j+1] = T.data[j]
			j = j - 1
		}
		T.data[j+1] = temp
	}
}

// ==========================================
// BAGIAN E: STATISTIK DATA
// ==========================================

// Menampilkan total penggunaan daya harian dan perangkat paling boros
func tampilkanStatistik(T TabPerangkat) {
	var i int
	var totalEnergi, konsumsiSaatIni, maxEnergi float64
	var namaBoros string

	totalEnergi = 0
	maxEnergi = -1

	for i = 0; i < T.n; i++ {
		konsumsiSaatIni = T.data[i].watt * T.data[i].durasi
		totalEnergi = totalEnergi + konsumsiSaatIni

		if konsumsiSaatIni > maxEnergi {
			maxEnergi = konsumsiSaatIni
			namaBoros = T.data[i].nama
		}
	}

	fmt.Printf("\n--- Statistik PowerLog ---\n")
	fmt.Printf("Total Penggunaan Daya Harian : %.2f Wh\n", totalEnergi)

	if T.n > 0 {
		fmt.Printf("Perangkat Paling Boros       : %s (Mengkonsumsi %.2f Wh/hari)\n", namaBoros, maxEnergi)
	} else {
		fmt.Println("Belum ada data perangkat untuk dianalisis.")
	}
}

// ==========================================
// FUNGSI UTAMA (MAIN INTERAKTIF)
// ==========================================
func main() {
	var daftar TabPerangkat
	var menu, indeks, indeksDitemukan int
	var nama, ruangan string
	var watt, durasi float64

	// Inisialisasi jumlah data awal
	daftar.n = 0

	// Pengisian data awal dummy (dari simulasi bawaan Anda)
	tambahPerangkat(&daftar, "Kulkas", "Dapur", 150.0, 24.0)    
	tambahPerangkat(&daftar, "AC", "Kamar", 400.0, 8.0)         
	tambahPerangkat(&daftar, "Televisi", "Keluarga", 80.0, 5.0) 
	tambahPerangkat(&daftar, "Blender", "Dapur", 250.0, 0.5)    

	for {
		fmt.Println("\n==========================================")
		fmt.Println("            MENU UTAMA POWERLOG           ")
		fmt.Println("==========================================")
		fmt.Println("1. Tampilkan Semua Perangkat")
		fmt.Println("2. Tambah Perangkat Baru")
		fmt.Println("3. Ubah Data Perangkat")
		fmt.Println("4. Hapus Perangkat")
		fmt.Println("5. Cari Perangkat Berdasarkan Ruangan (Sequential)")
		fmt.Println("6. Cari Perangkat Berdasarkan Nama (Binary)")
		fmt.Println("7. Urutkan Konsumsi Energi Tertinggi (Selection)")
		fmt.Println("8. Urutkan Abjad Nama Perangkat (Insertion)")
		fmt.Println("9. Tampilkan Statistik & Daftar Terboros")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&menu)

		if menu == 0 {
			fmt.Println("Keluar dari program. Terima kasih!")
			break
		}

		switch menu {
		case 1:
			fmt.Println("\n--- Daftar Perangkat Saat Ini ---")
			cetakSemuaPerangkat(daftar)
		case 2:
			fmt.Println("\n--- Tambah Perangkat Baru ---")
			fmt.Print("Nama Perangkat  : ")
			fmt.Scan(&nama)
			fmt.Print("Lokasi Ruangan  : ")
			fmt.Scan(&ruangan)
			fmt.Print("Daya (Watt)     : ")
			fmt.Scan(&watt)
			fmt.Print("Durasi (Jam)    : ")
			fmt.Scan(&durasi)
			tambahPerangkat(&daftar, nama, ruangan, watt, durasi)
		case 3:
			fmt.Println("\n--- Ubah Data Perangkat ---")
			cetakSemuaPerangkat(daftar)
			fmt.Print("Masukkan indeks data yang ingin diubah: ")
			fmt.Scan(&indeks)
			fmt.Print("Nama Baru       : ")
			fmt.Scan(&nama)
			fmt.Print("Ruangan Baru    : ")
			fmt.Scan(&ruangan)
			fmt.Print("Watt Baru       : ")
			fmt.Scan(&watt)
			fmt.Print("Durasi Baru     : ")
			fmt.Scan(&durasi)
			ubahPerangkat(&daftar, indeks, nama, ruangan, watt, durasi)
		case 4:
			fmt.Println("\n--- Hapus Perangkat ---")
			cetakSemuaPerangkat(daftar)
			fmt.Print("Masukkan indeks data yang ingin dihapus: ")
			fmt.Scan(&indeks)
			hapusPerangkat(&daftar, indeks)
		case 5:
			fmt.Println("\n--- Pencarian Berdasarkan Ruangan ---")
			fmt.Print("Masukkan nama ruangan yang dicari: ")
			fmt.Scan(&ruangan)
			cariBerdasarkanRuangan(daftar, ruangan)
		case 6:
			fmt.Println("\n--- Pencarian Berdasarkan Nama (Binary Search) ---")
			fmt.Print("Masukkan nama perangkat yang dicari: ")
			fmt.Scan(&nama)
			
			// FITUR TAMBAHAN: Otomatis mengurutkan abjad sebelum melakukan Binary Search
			urutAbjadNama(&daftar)
			indeksDitemukan = cariBerdasarkanNama(daftar, nama)
			
			if indeksDitemukan != -1 {
				fmt.Printf("Perangkat '%s' ditemukan pada indeks ke-%d.\n", nama, indeksDitemukan)
				fmt.Printf("Detail: [%s] %.1f Watt, %.1f jam/hari\n", 
					daftar.data[indeksDitemukan].ruangan, 
					daftar.data[indeksDitemukan].watt, 
					daftar.data[indeksDitemukan].durasi)
			} else {
				fmt.Println("Perangkat tidak ditemukan.")
			}
		case 7:
			urutEnergiTertinggi(&daftar)
			fmt.Println("✓ Data berhasil diurutkan berdasarkan konsumsi energi tertinggi.")
			cetakSemuaPerangkat(daftar)
		case 8:
			urutAbjadNama(&daftar)
			fmt.Println("✓ Data berhasil diurutkan berdasarkan abjad nama.")
			cetakSemuaPerangkat(daftar)
		case 9:
			// Menampilkan total energi harian
			tampilkanStatistik(daftar)
			
			// FITUR TAMBAHAN: Menampilkan "Daftar" peringkat perangkat paling boros keseluruhan
			fmt.Println("\n--- Peringkat Perangkat Dari yang Paling Boros ---")
			urutEnergiTertinggi(&daftar)
			cetakSemuaPerangkat(daftar)
		default:
			fmt.Println("Pilihan menu tidak valid. Silakan coba lagi.")
		}
	}
}