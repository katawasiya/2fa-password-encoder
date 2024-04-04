# HideYourStaff

This project is a tool for encrypting and decrypting files with using Google Authenticator and Secret, written in Go.
HideYourStuff will allow you to securely encrypt the file.

## Prerequisites

1. Open in browser `https://rapidapi.com/chdan/api/otp-authenticator`
2. Sign in using by github or gmail
3. Copy your X-RapidAPI-Key
4. Do `export RAPIDAPI_KEY="<your key>` in your terminal
      
## Usage 
`hideyourstuff encrypt <filename>`

### Using the Binary

1. Download the latest release.
2. Run the application.
3. Do `sudo mv hideyourstaff in your PATH` in your terminal

### Building from Source

1. `git clone git@github.com:katawasiya/hideyourstuff.git`
2. cd to directory with a project
3. `go build`
4. `sudo mv hideyourstaff in your PATH`

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

## License

This project is licensed under the MIT License. See the `LICENSE.md` file for details.
