package main

import "fmt"

const NMAX int = 100

type Perangkat struct {
	nama    string
	ruangan string
	watt    float64
	durasi  float64 // Lama pemakaian harian dalam jam
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
// Syarat: Data harus diurutkan berdasarkan nama (abjad) terlebih dahulu!
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
// FUNGSI UTAMA (MAIN)
// ==========================================
func main() {
	var daftar TabPerangkat
	var indeksDitemukan int

	// Inisialisasi jumlah data awal
	daftar.n = 0

	fmt.Println("=== SIMULASI TAMBAH DATA (Spesifikasi a & b) ===")
	tambahPerangkat(&daftar, "Kulkas", "Dapur", 150.0, 24.0)    // 3600 Wh
	tambahPerangkat(&daftar, "AC", "Kamar", 400.0, 8.0)         // 3200 Wh
	tambahPerangkat(&daftar, "Televisi", "Keluarga", 80.0, 5.0) // 400 Wh
	tambahPerangkat(&daftar, "Blender", "Dapur", 250.0, 0.5)    // 125 Wh

	fmt.Println("\n--- Daftar Perangkat Saat Ini ---")
	cetakSemuaPerangkat(daftar)

	fmt.Println("\n=== SIMULASI UBAH & HAPUS DATA (Spesifikasi a) ===")
	// Mengubah data Televisi (indeks 2) menjadi 100 Watt
	ubahPerangkat(&daftar, 2, "Televisi LED", "Keluarga", 100.0, 6.0)
	// Menghapus data Blender (indeks 3)
	hapusPerangkat(&daftar, 3)

	fmt.Println("\n--- Daftar Perangkat Setelah Perubahan ---")
	cetakSemuaPerangkat(daftar)

	// Menampilkan Statistik Awal
	tampilkanStatistik(daftar)

	fmt.Println("\n=== SIMULASI PENGURUTAN SELECTION SORT (Spesifikasi d) ===")
	fmt.Println("Mengurutkan berdasarkan konsumsi energi tertinggi:")
	urutEnergiTertinggi(&daftar)
	cetakSemuaPerangkat(daftar)

	fmt.Println("\n=== SIMULASI PENGURUTAN INSERTION SORT (Spesifikasi d) ===")
	fmt.Println("Mengurutkan berdasarkan abjad nama perangkat:")
	urutAbjadNama(&daftar)
	cetakSemuaPerangkat(daftar)

	fmt.Println("\n=== SIMULASI PENCARIAN (Spesifikasi c) ===")
	// 1. Sequential Search berdasarkan ruangan
	cariBerdasarkanRuangan(daftar, "Dapur")

	// 2. Binary Search berdasarkan nama
	// (Aman digunakan karena data sudah diurutkan berdasarkan abjad nama di langkah sebelumnya)
	indeksDitemukan = cariBerdasarkanNama(daftar, "AC")
	if indeksDitemukan != -1 {
		fmt.Printf("\nHasil Binary Search: Perangkat 'AC' ditemukan pada indeks ke-%d.\n", indeksDitemukan)
	} else {
		fmt.Println("\nHasil Binary Search: Perangkat tidak ditemukan.")
	}
}
