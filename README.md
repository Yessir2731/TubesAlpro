TUGAS BESAR ALGORITMA PEMOGRAMAN 2 (POWER LOG)

1. Pendahuluan 
Tugas Besar ini bertujuan untuk merancang dan membangun PowerLog, sebuah aplikasi pencatatan konsumsi listrik perangkat elektronik rumah tangga. Tujuan utama dari pengembangan aplikasi ini adalah untuk membantu pengguna memantau efisiensi energi dengan mencatat daya dalam satuan watt dan durasi penggunaan setiap perangkat. Melalui aplikasi ini, pengguna dapat mengetahui total konsumsi energi harian dalam satuan Watt-hour atau Wh serta mengidentifikasi perangkat mana yang paling boros energi.

2. Deskripsi Tubes 
Aplikasi PowerLog dikembangkan menggunakan bahasa pemrograman Golang dengan pendekatan prosedural, memanfaatkan tipe data bentukan struct dan struktur array statis dengan batas maksimal NMAX. Aplikasi ini berjalan berbasis Command Line Interface atau CLI interaktif. Fitur-fitur utama yang diimplementasikan meliputi:

Manipulasi Data CRUD: Pengguna dapat menambah, menampilkan, mengubah, dan menghapus data perangkat elektronik secara dinamis di dalam array.

Pencarian Searching: Menggunakan algoritma Sequential Search untuk mencari dan menampilkan daftar perangkat berdasarkan lokasi ruangan, serta Binary Search untuk mencari data perangkat spesifik berdasarkan namanya setelah data diurutkan.

Pengurutan Sorting: Menerapkan algoritma Selection Sort untuk mengurutkan perangkat secara descending berdasarkan konsumsi energi tertinggi, dan Insertion Sort untuk mengurutkan secara ascending berdasarkan abjad nama perangkat.

Statistik Penggunaan: Sistem secara otomatis mengalkulasi dan menampilkan total penggunaan daya harian serta mencari nilai maksimum untuk menemukan perangkat yang memakan daya paling besar.

3. Tantangan dan Solusi
Tantangan : Menyesuaikan penulisan sintaks program dengan standar dan batasan akademik yang ketat, khususnya larangan menggunakan deklarasi variabel singkat dan penanganan variabel bawaan. Selain itu, mengelola pergeseran indeks array secara manual saat fitur penghapusan data dijalankan membutuhkan ketelitian agar tidak terjadi kesalahan batas indeks.

Solusi : Mendisiplinkan penulisan kode dengan mendeklarasikan seluruh variabel secara eksplisit menggunakan kata kunci var beserta tipe datanya di awal setiap fungsi atau prosedur. Untuk menangani penghapusan data, diimplementasikan logika iterasi yang menggeser elemen array ke kiri secara terstruktur, dan memastikan nilai batas data selalu dikurangi dengan tepat.

4. Kesimpulan dan Rekomendasi
Kesimpulan: Aplikasi PowerLog berhasil memenuhi seluruh spesifikasi yang diminta. Logika dasar algoritma seperti pencarian dan pengurutan telah berhasil diimplementasikan dari awal tanpa bergantung pada pustaka bawaan Golang. Program mampu merangkum statistik penggunaan energi 
secara akurat melalui antarmuka teks yang terstruktur.

Rekomendasi: Untuk pengembangan lebih lanjut, disarankan untuk mengimplementasikan fitur pemrosesan file seperti pembacaan dan penulisan ke file teks. Saat ini, data array masih bersifat sementara di dalam memori, sehingga data akan hilang setiap kali program ditutup. Menyimpan data ke file akan membuat aplikasi ini jauh lebih fungsional untuk penggunaan sehari-hari.

5. Referensi
- Modul  Pembelajaran Algoritma Pemrograman 2 di LMS.
- Modul Praktikum Algoritma Pemrograman 2. 
