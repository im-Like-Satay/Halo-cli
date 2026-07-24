# Halo 
<img width=300 src=image/halo.png alt="Halo Logo" />

"Halo" adalah sebuah ai chat berbasis CLI sederhana dan cepat, serta mendukung penuh ekosistem Gemini AI. Ditulis menggunakan bahasa Go.

**English version** [here](README-EN.md)

## Filosofi
Filosofi dari "Halo" sendiri adalah kesederhanaan.
- tidak ada konfigurasi rumit.
- tidak ada biaya sepeserpun.
- tidak ada flag rumit yang perlu dihafal atau dipelajari terlebih dahulu.
- cukup paste apikey gemini anda lalu gunakan "Halo" sesuka hati ❤️.

## Instalasi Cepat
1. Clone repo "Halo".
```bash
git clone https://github.com/im-Like-Satay/Halo-cli 

cd Halo-cli 
```
2. Konfigurasi gemini apikey di `Halo-cli/internal/ai.go` pada baris 16, paste apikey Gemini anda di bagian ini. Ganti <paste_apikey_here> dengan api Gemini anda.
```go
  APIKey: "<paste_apikey_here>",
```
3. Install dependensi
```bash
go mod tidy
```

4. Build Halo
```bash
# Windows
go build -o halo.exe
# Linux
go build -o halo
```
Sekarang "Halo" siap digunakan
## Penggunaan
Cukup gunakan seperti ini.
```bash
halo "apa itu golang"
```

## Screenshot
<img width=500 src="image/ss-one.png" alt="Glamour example">
<img width=500 src="image/ss-two.png" alt="Glamour example">

## TODO
- konfigurasi yang lebih mudah.
- instalasi yang lebih mudah.

## Lisensi
MIT
