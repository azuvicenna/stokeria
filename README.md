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

Semua endpoint API diawali dengan *prefix* `/api`.

> **Catatan Autentikasi Transaksi**: Untuk endpoint `POST` pada transaksi (*Inbound*, *Outbound*, dan *Stock Adjustment*), sertakan header `X-User-Id: <UUID-USER>` sebagai identitas staff/user yang melakukan transaksi.

### 1. Sistem & Health Check
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `GET` | `/api/health` | Mengecek status kesehatan service aplikasi |

---

### 2. Kategori Produk (Categories)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/categories` | Membuat kategori produk baru |
| `GET` | `/api/categories` | Mengambil seluruh daftar kategori produk |
| `GET` | `/api/categories/:id` | Mengambil detail kategori berdasarkan ID |
| `PUT` | `/api/categories/:id` | Memperbarui nama/deskripsi kategori berdasarkan ID |
| `DELETE` | `/api/categories/:id` | Menghapus data kategori berdasarkan ID |

---

### 3. Pengguna & Autentikasi (Users & Auth)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/login` | Otentikasi dan login user dengan email dan password |
| `POST` | `/api/users` | Mendaftarkan pengguna baru (Role: `admin` / `warehouse_staff`) |
| `GET` | `/api/users` | Mengambil seluruh daftar pengguna |
| `GET` | `/api/users/:id` | Mengambil detail pengguna berdasarkan ID |
| `PUT` | `/api/users/:id` | Memperbarui informasi profil pengguna berdasarkan ID |
| `DELETE` | `/api/users/:id` | Menghapus data pengguna berdasarkan ID |

---

### 4. Pemasok (Suppliers)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/suppliers` | Menambahkan data pemasok (*supplier*) baru |
| `GET` | `/api/suppliers` | Mengambil seluruh daftar pemasok |
| `GET` | `/api/suppliers/:id` | Mengambil detail pemasok berdasarkan ID |
| `PUT` | `/api/suppliers/:id` | Memperbarui data pemasok berdasarkan ID |
| `DELETE` | `/api/suppliers/:id` | Menghapus data pemasok berdasarkan ID |

---

### 5. Gudang (Warehouses)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/warehouses` | Menambahkan data gudang baru (kode unik & alamat) |
| `GET` | `/api/warehouses` | Mengambil seluruh daftar gudang |
| `GET` | `/api/warehouses/:id` | Mengambil detail gudang berdasarkan ID |
| `PUT` | `/api/warehouses/:id` | Memperbarui informasi gudang berdasarkan ID |
| `DELETE` | `/api/warehouses/:id` | Menghapus data gudang berdasarkan ID |

---

### 6. Produk (Products)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/products` | Menambahkan produk baru (SKU unik, harga, dan relasi kategori) |
| `GET` | `/api/products` | Mengambil seluruh daftar produk beserta info kategorinya |
| `GET` | `/api/products/:id` | Mengambil detail produk beserta info kategorinya |
| `PUT` | `/api/products/:id` | Memperbarui data produk berdasarkan ID |
| `DELETE` | `/api/products/:id` | Menghapus produk berdasarkan ID |

---

### 7. Stok Gudang (Inventories)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `GET` | `/api/inventories` | Mengambil seluruh daftar stok inventory di semua gudang |
| `GET` | `/api/inventories/:id` | Mengambil detail stok inventory berdasarkan ID |

---

### 8. Transaksi Masuk (Inbound Transactions)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/inbound-transactions` | Mencatat barang masuk dari supplier ke gudang dan **menambah stok inventory** otomatis |
| `GET` | `/api/inbound-transactions` | Mengambil riwayat seluruh transaksi masuk |
| `GET` | `/api/inbound-transactions/:id` | Mengambil detail transaksi masuk dan rincian item produknya |

---

### 9. Transaksi Keluar (Outbound Transactions)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/outbound-transactions` | Mencatat barang keluar ke destinasi tujuan dan **memotong stok inventory** (validasi kecukupan stok) |
| `GET` | `/api/outbound-transactions` | Mengambil riwayat seluruh transaksi keluar |
| `GET` | `/api/outbound-transactions/:id` | Mengambil detail transaksi keluar dan rincian item produknya |

---

### 10. Penyesuaian Stok (Stock Adjustments)
| Method | Endpoint | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/stock-adjustments` | Mencatat stok opname / barang rusak dan **memperbarui kuantitas stok gudang** |
| `GET` | `/api/stock-adjustments` | Mengambil riwayat seluruh penyesuaian stok |
| `GET` | `/api/stock-adjustments/:id` | Mengambil detail riwayat penyesuaian stok berdasarkan ID |

---

## Format Respons Standar

### Respons Sukses
```json
{
  "code": 200,
  "status": "OK",
  "data": { ... }
}
```

### Respons Error (Ditangani oleh `exception.ErrorHandler`)
```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "errors": {
    "Email": "Validation failed on tag: email"
  }
}
```
