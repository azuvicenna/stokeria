# Stokeria API

Stokeria adalah sistem backend manajemen inventaris, pergudangan, dan transaksi stok barang (Inbound, Outbound, dan Stock Adjustment) yang dibangun menggunakan Go, Fiber v2, dan GORM MySQL.

---

## Perintah yang Digunakan (Commands)

Berikut adalah daftar perintah yang digunakan dalam proyek ini, dikelompokkan berdasarkan fungsinya:

### 1. Instalasi Dependensi (Dependencies)
```bash
# Web Framework
go get github.com/gofiber/fiber/v2

# Validator Request
go get github.com/go-playground/validator/v10

# Logging & Konfigurasi
go get github.com/sirupsen/logrus
go get github.com/spf13/viper

# Database ORM & Driver MySQL
go get gorm.io/gorm
go get gorm.io/driver/mysql

# Database Migration
go get github.com/golang-migrate/migrate/v4

# JSON Web Token (JWT)
go get github.com/golang-jwt/jwt/v5

# UUID Generator & Testing
go get github.com/google/uuid
go get github.com/stretchr/testify
```

### 2. Sinkronisasi Modul
```bash
go mod tidy
```

### 3. Database Seeder
```bash
# Menjalankan seeder untuk mengisi data dummy awal
go run main.go -seed
```

### 4. Menjalankan Server Aplikasi
```bash
# Menjalankan server aplikasi (default port :3000)
go run main.go
```

### 5. Menjalankan Pengujian (Testing)
```bash
# Menjalankan seluruh automated unit/integration test
go test -v ./...
```

---

## Daftar Endpoint API

Semua endpoint API diawali dengan prefix `/api`.

### 1. Sistem & Health Check (Publik)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `GET` | `/api/health` | Mengecek status kesehatan service aplikasi |

---

### 2. Autentikasi (Publik)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/login` | Otentikasi user & mengembalikan Bearer JWT Token beserta data profil user |

---

### 3. Pengguna
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/users` | Mendaftarkan pengguna baru (Role: `admin` / `warehouse_staff`) |
| `GET` | `/api/users` | Mengambil seluruh daftar pengguna |
| `GET` | `/api/users/:id` | Mengambil detail pengguna berdasarkan ID |
| `PUT` | `/api/users/:id` | Memperbarui informasi profil pengguna berdasarkan ID |
| `DELETE` | `/api/users/:id` | Menghapus data pengguna berdasarkan ID |

---

### 4. Kategori Produk
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/categories` | Membuat kategori produk baru |
| `GET` | `/api/categories` | Mengambil seluruh daftar kategori produk |
| `GET` | `/api/categories/:id` | Mengambil detail kategori berdasarkan ID |
| `PUT` | `/api/categories/:id` | Memperbarui nama/deskripsi kategori berdasarkan ID |
| `DELETE` | `/api/categories/:id` | Menghapus data kategori berdasarkan ID |

---

### 5. Pemasok
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/suppliers` | Menambahkan data pemasok (supplier) baru |
| `GET` | `/api/suppliers` | Mengambil seluruh daftar pemasok |
| `GET` | `/api/suppliers/:id` | Mengambil detail pemasok berdasarkan ID |
| `PUT` | `/api/suppliers/:id` | Memperbarui data pemasok berdasarkan ID |
| `DELETE` | `/api/suppliers/:id` | Menghapus data pemasok berdasarkan ID |

---

### 6. Gudang
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/warehouses` | Menambahkan data gudang baru (kode unik & alamat) |
| `GET` | `/api/warehouses` | Mengambil seluruh daftar gudang |
| `GET` | `/api/warehouses/:id` | Mengambil detail gudang berdasarkan ID |
| `PUT` | `/api/warehouses/:id` | Memperbarui informasi gudang berdasarkan ID |
| `DELETE` | `/api/warehouses/:id` | Menghapus data gudang berdasarkan ID |

---

### 7. Produk
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/products` | Menambahkan produk baru (SKU unik, harga, dan relasi kategori) |
| `GET` | `/api/products` | Mengambil seluruh daftar produk beserta info kategorinya |
| `GET` | `/api/products/:id` | Mengambil detail produk beserta info kategorinya |
| `PUT` | `/api/products/:id` | Memperbarui data produk berdasarkan ID |
| `DELETE` | `/api/products/:id` | Menghapus produk berdasarkan ID |

---

### 8. Stok Gudang
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `GET` | `/api/inventories` | Mengambil seluruh daftar stok inventory di semua gudang |
| `GET` | `/api/inventories/:id` | Mengambil detail stok inventory berdasarkan ID |

---

### 9. Transaksi Masuk
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/inbound-transactions` | Mencatat barang masuk dari supplier ke gudang dan menambah stok inventory otomatis |
| `GET` | `/api/inbound-transactions` | Mengambil riwayat seluruh transaksi masuk |
| `GET` | `/api/inbound-transactions/:id` | Mengambil detail transaksi masuk dan rincian item produknya |

---

### 10. Transaksi Keluar
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/outbound-transactions` | Mencatat barang keluar ke destinasi tujuan dan memotong stok inventory (validasi kecukupan stok) |
| `GET` | `/api/outbound-transactions` | Mengambil riwayat seluruh transaksi keluar |
| `GET` | `/api/outbound-transactions/:id` | Mengambil detail transaksi keluar dan rincian item produknya |

---

### 11. Penyesuaian Stok
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/stock-adjustments` | Mencatat stok opname / barang rusak dan memperbarui kuantitas stok gudang |
| `GET` | `/api/stock-adjustments` | Mengambil riwayat seluruh penyesuaian stok |
| `GET` | `/api/stock-adjustments/:id` | Mengambil detail riwayat penyesuaian stok berdasarkan ID |

---

## Format Respons Standar

### Respons Sukses Login (`POST /api/login`)
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "user": {
      "id": "e30d17d5-d0c3-4d45-926c-d2ff86e74b59",
      "email": "user1@stokeria.com",
      "name": "User 1",
      "role": "admin",
      "created_at": "2026-11-03T06:00:00Z",
      "updated_at": "2026-11-03T06:00:00Z"
    }
  }
}
```

### Respons Error Terpusat (`exception.ErrorHandler`)
```json
{
  "code": 401,
  "status": "UNAUTHORIZED",
  "errors": {
    "message": "Missing Authorization header"
  }
}
```
