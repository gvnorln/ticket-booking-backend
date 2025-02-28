# Tahap 1: Build aplikasi menggunakan image Go
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Salin file go.mod dan go.sum untuk caching dependency
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh source code ke dalam container
COPY . .

# Build aplikasi; sesuaikan path package sesuai dengan struktur project Anda
RUN go build -o ticket-booking-backend ./main.go

# Tahap 2: Buat image final yang ringan menggunakan Alpine
FROM alpine:latest
WORKDIR /app

# Salin file binary yang sudah di-build dari tahap builder
COPY --from=builder /app/ticket-booking-backend .

# Ekspos port yang digunakan aplikasi (misalnya 8080)
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["./ticket-booking-backend"]
