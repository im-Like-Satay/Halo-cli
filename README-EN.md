# Halo 
<img width=300 src=image/halo.png alt="Halo Logo" />

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/AI-Gemini_3.5-8E7CC3?style=for-the-badge&logo=googlegemini&logoColor=white" />
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" />
</p>


[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![Gemini AI](https://img.shields.io/badge/Gemini-3.5_Flash--Lite-8E7CC3?style=flat-square&logo=googlegemini)](https://aistudio.google.com)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
"Halo" is a simple and fast CLI-based AI chat, and fully supports the Gemini AI ecosystem. Written in Go.

**English version** [here](README-EN.md)

## Philosophy
The philosophy of "Halo" itself is simplicity.
- no complicated configuration.
- no cost whatsoever.
- no complex flags that need to be memorized or learned first.
- just paste your Gemini API key and use "Halo" as you like ❤️.

## Quick Installation
```bash
git clone https://github.com/im-Like-Satay/Halo-cli
cd Halo-cli 
go mod tidy

# Windows
go build -o halo.exe .

# Linux / macOS
go build -o halo .

# Set API Key 
halo set <apikey>
```

**Or Download Binary [Here](https://github.com/im-Like-Satay/Halo-cli/releases)**

## Usage
```bash
halo "apa itu golang"
```

## Screenshot
<img width=500 src="image/ss-one.png" alt="Glamour example">
<img width=500 src="image/ss-two.png" alt="Glamour example">

## TODO
[✅] easier configuration.
[✅] easier installation.

## License
MIT
