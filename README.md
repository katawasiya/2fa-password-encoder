# File Encryptor

This project is a tool for encrypting and decrypting files with using Google Authenticator and Secret, written in Go.
HideYourStuff will allow you to securely encrypt the file.

## Project Structure

- `.github/workflows/go.yml`
- `commands/commands.go`
- `example_encrypted.csv`
- `example.csv`
- `go.mod`
- `go.sum`
- `main.go`
- `pckg/2fa/2fa.go`
- `pckg/crypt/crypt.go`
- `pckg/qr/qr.go`
- `pckg/secrets/secret_key.go`

## Usage 

1. Open in browser https://rapidapi.com/chdan/api/otp-authenticator
2. Sign in using github/gmail
3. Copy your X-RapidAPI-Key
4. In terminal do: export RAPIDAPI_KEY="<your key>" 
5. In terminal do: hideyourstuff <encrypt/decrypt> <filename>
6. Save your secret 
7. Scan your QR Code with Google Authenticator

### Using the Binary

1. Download the latest release.
2. Run the application.

### Building from Source

1. git clone git@github.com:katawasiya/hideyourstuff.git
2. cd to directory with a project
3. go build
4. sudo mv hideyourstaff in your PATH

## License

This project is licensed under the MIT License. See the `LICENSE.md` file for details.