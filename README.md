# Implementasi Saga Pattern dengan gRPC

## Deskripsi Proyek
Proyek ini adalah implementasi Saga Pattern menggunakan gRPC untuk menangani transaksi terdistribusi dengan tiga microservice: Order, Payment, dan Shipping.

## Prasyarat
- Go (versi 1.21 atau lebih baru)
- Protocol Buffers Compiler (protoc)
- gRPC

## Persiapan Awal
1. Clone repository
```bash
git clone https://github.com/122140015devaahmad/122140015_Evaluasi1Tugas2.git
cd grpc-saga
```

2. Instal dependensi
```bash
go mod tidy
```

## Menjalankan Services

### Kompilasi Proto Files
```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

### Jalankan Services (dalam terminal terpisah)
1. Order Service
```bash
go run order_service/main.go
```

2. Payment Service
```bash
go run payment_service/main.go
```

3. Shipping Service
```bash
go run shipping_service/main.go
```

4. Saga Orchestrator
```bash
go run saga_orchestrator/main.go
```

## Skenario Pengujian
- Berhasil: Semua service berjalan tanpa masalah
- Gagal di Order Service
- Gagal di Payment Service
- Gagal di Shipping Service

## Konsep Utama
- Saga Pattern untuk transaksi terdistribusi
- Kompensasi otomatis jika salah satu service gagal
- Komunikasi antar service menggunakan gRPC

## Kontributor
Deva Ahmad (122140015)
