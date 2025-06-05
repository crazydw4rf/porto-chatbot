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
   go build -o porto-chatbot ./cmd/porto-chatbot/main.go
   ```

5. **Jalankan aplikasi**
   ```bash
   ./porto-chatbot
   ```
   Secara default, aplikasi akan berjalan di port 3000.

## Endpoint

- `POST /chat`
  Contoh request payload:
  ```json
  {
    "question": "Apa pengalaman Ucup dengan Docker?"
  }
  ```
  Contoh Response:
  ```json
  {
    "answer": "Ucup aktif dalam proses deployment menggunakan Docker dan CI/CD pipeline dengan GitHub Actions..."
  }
  ```

## Deploy ke Vercel Serverless Function

Proyek ini sudah dikonfigurasi untuk bisa dideploy ke Vercel sebagai serverless function. Ikuti langkah berikut:

### 1. Set Environment Variables di Vercel

Setelah membuat project di Vercel:

1. Buka dashboard Vercel project kamu
2. Pergi ke **Settings** → **Environment Variables**
3. Tambahkan variabel berikut:
   - `GEMINI_API_KEY`: API key Gemini kamu
   - `CORS_DOMAINS`: Domain yang diizinkan (contoh: `https://yourapp.vercel.app` atau `*` untuk semua)

### 2. Deploy

Ada dua cara untuk deploy:

**Cara 1: Melalui GitHub (Recommended)**
1. Push project ke GitHub repository
2. Import repository di Vercel dashboard
3. Vercel akan otomatis build dan deploy

**Cara 2: Melalui Vercel CLI**
```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel

# Follow the prompts
```

### 3. Testing

Setelah deploy berhasil, kamu bisa test endpoint:
```bash
curl -X POST https://your-app.vercel.app/chat \
  -H "Content-Type: application/json" \
  -d '{"question": "Apa pengalaman Ucup dengan Docker?"}'
```

## Catatan

- Folder `web/` dapat digunakan untuk menaruh file frontend (misal: index.html).
- Jika ingin mengubah port, edit bagian berikut di `main.go`:
  ```go
  err := app.Listen(":3000")
  ```
  Ganti `3000` dengan port yang diinginkan.

## Mengubah Instruksi Portfolio

Jika ingin mengganti konteks portfolio yang digunakan chatbot, ubah nilai konstanta `UcupPortfolio` pada file `instruction/portfolio.go`.

Contoh:
```go
const UcupPortfolio = `Berikut adalah portfolio singkat tentang Budi:
Budi adalah seorang Backend Developer yang berpengalaman dalam membangun API dengan Go dan PostgreSQL...`
```
Pastikan untuk menyesuaikan isi string tersebut sesuai kebutuhan.
