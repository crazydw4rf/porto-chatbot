# porto-chatbot

## Cara Build & Menjalankan Aplikasi

1. **Clone repository ini**
   ```bash
   git clone github.com/crazydw4rf/porto-chatbot.git
   cd porto-chatbot
   ```

2. **Install dependencies**
   Pastikan kamu sudah menginstall Go (minimal versi 1.20).
   Lalu jalankan:
   ```bash
   go mod download
   ```

3. **Konfigurasi file .env**
   Salin file `.env.example` menjadi `.env` di root folder:
   ```bash
   cp .env.example .env
   ```
    Kemudian, buka file `.env` dan masukkan API key Gemini yang valid pada variabel `GEMINI_API_KEY`.

4. **Build aplikasi**
   ```bash
   go build -o porto-chatbot ./cmd/porto-chatbot
   ```

5. **Jalankan aplikasi**
   ```bash
   ./porto-chatbot
   ```
   Secara default, aplikasi akan berjalan di port 3000.

## Variabel Environment yang Perlu Diubah

Variabel-variabel berikut dibaca menggunakan Viper, baik dari file `.env` maupun environment variables sistem:

- `GEMINI_API_KEY`
  Masukkan API key Gemini yang valid agar chatbot dapat berfungsi.

## Endpoint

- `POST /chat`
  Kirim pertanyaan tentang portfolio Ucup dalam format JSON:
  ```json
  {
    "question": "Apa pengalaman Ucup dengan Docker?"
  }
  ```
  Response:
  ```json
  {
    "answer": "Ucup aktif dalam proses deployment menggunakan Docker dan CI/CD pipeline dengan GitHub Actions..."
  }
  ```

## Catatan

- Folder `web/` dapat digunakan untuk menaruh file frontend (misal: index.html).
- Jika ingin mengubah port, edit bagian berikut di `main.go`:
  ```go
  err := app.Listen(":3000")
  ```
  Ganti `3000` dengan port yang diinginkan.

## Mengubah portfolioContext

Jika ingin mengganti konteks portfolio yang digunakan chatbot, ubah nilai variabel `portfolioContext` pada file `cmd/porto-chatbot/main.go`.

Contoh:
```go
const portfolioContext = `Berikut adalah portfolio singkat tentang Budi:
Budi adalah seorang Backend Developer yang berpengalaman dalam membangun API dengan Go dan PostgreSQL...`
```
Pastikan untuk menyesuaikan isi string tersebut sesuai kebutuhan Anda.
