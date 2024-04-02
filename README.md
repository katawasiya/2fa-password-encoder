# File Encryptor

This project is a tool for encrypting and decrypting files, written in Go.

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

1. Download the latest release.
2. Run the application.

## Local Build

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run `go build` to compile the project.
4. Run the compiled binary with the necessary arguments to encrypt or decrypt a file.

## Functions

- `crypt.Encrypt(filename string)`: Encrypts the file with the given name. Returns an `EncryptStruct` and an error if any.
- `crypt.Decrypt(secret string, filename string)`: Decrypts the file with the given name using the given secret key. Returns an error if any.

## Testing

To run tests, use the `go test ./...` command in the root directory of the project.

## CI/CD

The project uses GitHub Actions for automatic testing of the code on every push or pull request. See [`.github/workflows/go.yml`](command:_github.copilot.openRelativePath?%5B%22.github%2Fworkflows%2Fgo.yml%22%5D ".github/workflows/go.yml") for details.

## License

This project is licensed under the MIT License. See the `LICENSE.md` file for details.
