# app-movie-festival

Aplikasi **Movie Festival** adalah sistem manajemen film yang digunakan pada sebuah festival. Aplikasi ini memiliki fitur-fitur berikut:

## Fitur Utama

1. **Admin Management:**
   - Sebelum festival dimulai, admin dapat menambahkan data film ke dalam sistem, termasuk mengunggah file film ke database dan local storage.
2. **Pengunjung Aplikasi:**
   - Selama festival berlangsung:
     - Pengunjung dapat mencari film yang ingin mereka tonton.
     - Pengunjung dapat memberikan vote pada film yang mereka sukai.
     - Pengunjung dapat membatalkan vote melalui daftar history vote.
3. **Laporan Pasca-Festival:**
   - Setelah festival selesai, admin dapat melihat:
     - Judul film dengan jumlah vote terbanyak.
     - Genre film dengan jumlah view terbanyak.

**Catatan:** Pengunjung harus login untuk memberikan vote pada film.

---

## Cara Menjalankan Aplikasi

### 1. Persiapan Awal

#### a. **Install Golang**

- Pastikan **Golang** terinstal di sistem Anda. Jalankan perintah berikut:
  ```bash
  brew install go
  ```
- Verifikasi instalasi dengan perintah:
  ```bash
  go version
  ```

#### b. **Install MySQL**

- Instal **MySQL** dengan **Homebrew**:
  ```bash
  brew install mysql
  ```
- Mulai layanan MySQL:
  ```bash
  brew services start mysql
  ```

---

### 2. Membuat Database

1. Masuk ke MySQL CLI:
   ```bash
   mysql -u root -p
   ```
   Masukkan password root Anda saat diminta.
2. Buat database dengan nama `parcel`:
   ```sql
   CREATE DATABASE parcel;
   ```
3. Keluar dari MySQL:
   ```sql
   EXIT;
   ```

---

### 3. Clone Repository

Clone repository ke lokal Anda:

```bash
git clone https://github.com/Gading09/app-movie-festival.git
cd app-movie-festival
```

---

### 4. Instalasi Dependensi

Pastikan semua dependensi yang diperlukan terinstal dengan menjalankan:

```bash
go mod tidy
```

---

### 5. Menjalankan Aplikasi

Jalankan aplikasi menggunakan perintah:

```bash
go run ./cmd/main.go
```

---

### Struktur Proyek:

- `/cmd`: Berisi file untuk menjalankan aplikasi.
- `/config`: Untuk mengambil data konfigurasi dari file .env.
- `/database`: Menangani koneksi ke database.
- `/delivery`: Menangani pengiriman data melalui berbagai protokol (seperti HTTP atau gRPC).
- `/domain`: Kumpulan paket yang berisi beberapa layer dan file handler.
- `/feature`: Logika bisnis aplikasi.
- `/repository`: Mengelola interaksi langsung dengan database.
- `/model`: Kumpulan struct yang digunakan untuk mendefinisikan struktur data di database.

## Teknologi yang Digunakan:

- **Go**: Bahasa pemrograman utama.
- **GORM**: ORM untuk interaksi dengan database MySQL.
- **MySQL**: Database yang digunakan untuk menyimpan data.
- **JWT**: Untuk autentikasi dan otorisasi pengguna.

---
