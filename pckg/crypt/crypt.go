package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	secretKey "encryptor/pckg/secrets"
	"fmt"
	"io"
	"os"

	"github.com/pquerna/otp/totp"
)

type EncryptScruct struct {
	Secret   *secretKey.SecretKeyStruct
	Filename string
}

type DecryptScruct struct {
	Secret   string
	Filename string
}

// this func encrypts the file by using the secret key
func Encrypt(filename string) (*EncryptScruct, error) {

	s := &EncryptScruct{
		Secret:   secretKey.InitializeSecret(),
		Filename: filename,
	}

	content, err := os.ReadFile(s.Filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(s.Secret.SecretKey))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(content))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], content)

	err = os.WriteFile(filename, ciphertext, 0644)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// this func decrypts the file by using the secret key and the otp code
func Decrypt(decryptionKey string, filename string, otpCode string) (*DecryptScruct, error) {

	s := &DecryptScruct{
		Secret:   decryptionKey,
		Filename: filename,
	}

	valid := totp.Validate(otpCode, s.Secret)
	if !valid {
		fmt.Println("Invalid OTP code")
		return nil, fmt.Errorf("invalid OTP code")
	}

	content, err := os.ReadFile(s.Filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	block, err := aes.NewCipher([]byte(s.Secret))
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return nil, err
	}

	iv := content[:aes.BlockSize]
	content = content[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(content, content)

	// Write out the decrypted content.
	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return nil, err
	}

	fmt.Println("Decryption complete")
	return nil, err
}
