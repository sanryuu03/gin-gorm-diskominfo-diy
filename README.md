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