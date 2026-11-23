# gin-gorm-diskominfo-diy

API booking ruangan menggunakan Go, Gin, dan GORM (PostgreSQL).

---

## Menjalankan dengan Docker

### Prasyarat

- Docker Engine + Docker Compose Plugin (`docker compose version`)
- File `.env` di direktori utama repositori

### Konfigurasi `.env`

Docker Compose membaca `.env` dari host untuk substitusi variabel. File ini **tidak ikut masuk image** karena sudah terdaftar di `.dockerignore`.

```env
DB_USER=postgres
DB_PASSWORD=passwordku
DB_NAME=diskominfo_diy
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Jakarta
APP_PORT=8080
```

> `DB_HOST` tidak perlu diisi untuk Docker. Compose otomatis mengatur `DB_HOST=db` di dalam container. Nilai `DB_PORT` di `.env` dipakai untuk port mapping host, sedangkan container tetap memakai `5432`.

### Menjalankan aplikasi

```bash
# Build image + jalankan Postgres dan aplikasi
docker compose up -d --build

# Lihat log aplikasi
docker compose logs -f app

# Cek health check
curl -i http://localhost:8080/ping
```

Saat start, service `app` menjalankan `./app -migrate` terlebih dahulu lalu menyalakan server Gin di port `8080`.

### Seed data

Seed **tidak** dijalankan otomatis karena belum idempoten (email unik akan bertabrakan jika dijalankan dua kali).

```bash
docker compose run --rm seed
```

Mengulang dari data kosong:

```bash
docker compose run --rm truncate
docker compose run --rm seed
```

### Service yang tersedia

| Service | Profil | Fungsi |
|---|---|---|
| `db` | default | PostgreSQL 16 + volume `pgdata` |
| `app` | default | API server (auto-migrate saat start) |
| `migrate` | `tools` | Menjalankan migrasi manual |
| `seed` | `tools` | Mengisi data users, rooms, bookings |
| `truncate` | `tools` | Mengosongkan seluruh tabel |

Service dengan profil `tools` hanya jalan lewat `docker compose run`, bukan saat `docker compose up`:

```bash
docker compose run --rm migrate
```

### Perintah berguna

```bash
docker compose ps                    # status container
docker compose restart app           # restart API saja
docker compose down                  # stop, data DB tetap tersimpan
docker compose down -v               # stop + hapus volume DB
docker compose up -d --build --force-recreate  # rebuild dari nol
```

Jika port `8080` bentrok, ubah `APP_PORT` di `.env`, lalu akses API lewat port tersebut. `BASE_URL` di bagian curl di bawah harus disesuaikan juga.

### Catatan build

`Dockerfile` memakai dua stage: `golang:1.26-alpine` untuk kompilasi (`CGO_ENABLED=0`) dan `alpine:3.21` untuk runtime. Hasilnya image final hanya berisi binary `app` plus `ca-certificates` dan `tzdata`.

---

## Konvensi yang dipakai

```bash
export BASE_URL=http://localhost:8080
```

---

## `GET /ping`

Health check sederhana.

```bash
curl -i "$BASE_URL/ping"
```

Response `200 OK`:

```json
{"message":"pong"}
```

---

## `POST /api/v1/auth/login`

Login dengan email + password, mengembalikan JWT.

### Berhasil — `200 OK`

```bash
curl -i -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user1@test.com",
    "password": "password123"
  }'
```

Response:

```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
```

### Email tidak ditemukan — `401 Unauthorized`

```bash
curl -i -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "tidakada@test.com",
    "password": "password123"
  }'
```

### Password salah — `401 Unauthorized`

```bash
curl -i -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user1@test.com",
    "password": "salahpassword"
  }'
```

### Simpan token untuk endpoint terproteksi

```bash
TOKEN=$(curl -s -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"user1@test.com","password":"password123"}' \
  | jq -r '.token')

echo "$TOKEN"
```

Pemakaiannya:

```bash
curl -i "$BASE_URL/api/v1/rooms" \
  -H "Authorization: Bearer $TOKEN"
```

---