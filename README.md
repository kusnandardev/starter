# How to use this golang starter

## Prequisite

- Golang v1.12.X
- Docker & dokcer-compose
- Fresh (golang autoreload)
  > go get -u github.com/pilu/fresh
- swago (swagger generator)
  > go get -u github.com/swaggo/swag/cmd/swag

## How To start

- jalankan docker-compose
  > docker-compose up
- generate dokumentasi terakhir
  > swag init
- jalankan project
  > fresh

## Note

- semua config ada di `/conf/app.ini` (next pindah ke env)
- semua yang berhubungan dengan query diletakan di `model`
- data dari model dimanipulasi di `service`
- handler untuk request/respons diletakan di `router` 