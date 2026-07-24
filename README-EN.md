# Halo 
<img width=300 src=image/halo.png alt="Halo Logo" />

"Halo" is a simple and fast CLI-based AI chat that fully supports the Gemini AI ecosystem. Written in Go.

**Indonesia version** [here](README-EN.md)

## Philosophy
The philosophy of "Halo" itself is simplicity.
- no complicated configuration.
- no cost whatsoever.
- no complex flags that need to be memorized or learned first.
- just paste your Gemini API key and use "Halo" as you like ❤️.

## Quick Installation
1. Clone the "Halo" repo.
```bash
git clone https://github.com/im-Like-Satay/Halo-cli 

cd Halo-cli 
```
2. Configure the Gemini API key in `Halo-cli/internal/ai.go` on line 16, paste your Gemini API key in this section. Replace <paste_apikey_here> with your Gemini API.
```go
  APIKey: "<paste_apikey_here>",
```
3. Install dependencies
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
Now "Halo" is ready to use
## Usage
Just use it like this.
```bash
halo "apa itu golang"
```

## Screenshot
<img width=500 src="image/ss-one.png" alt="Glamour example">
<img width=500 src="image/ss-two.png" alt="Glamour example">

## TODO
- very simple config.
- very simple install.

## License
MIT
