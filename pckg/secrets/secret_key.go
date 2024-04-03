package secretKey

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type SecretKeyStruct struct {
	SecretKey    string
	SecretLength int
	QRCode       string
}

type SecretKeyInterface interface {
	RandomSecretLength() *SecretKeyStruct
	GenerateSecretKey() *SecretKeyStruct
	InitializeSecret() *SecretKeyStruct
	CreateQRCode() *SecretKeyStruct
}

func (s *SecretKeyStruct) RandomSecretLength() *SecretKeyStruct {
	validLengths := []int{16, 24, 32}
	s.SecretLength = validLengths[rand.Intn(len(validLengths))]

	return s
}

func (s *SecretKeyStruct) GenerateSecretKey() string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567" // Base32 set
	b := make([]byte, s.SecretLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	s.SecretKey = string(b)
	return s.SecretKey
}

func (s *SecretKeyStruct) CreateQRCode() *SecretKeyStruct {
	url := "https://otp-authenticator.p.rapidapi.com/qr2/?data=otpauth%3A%2F%2Ftotp%2FHomeCorp%3AUser1%3Fsecret%3D7UPL3UKHTKUHYY2O%26issuer%3DHomeCorp&size=5&level=S"
	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "5fa16887e7msh10ce850bbebe7a7p10ecf5jsnbacaae827963")
	req.Header.Add("X-RapidAPI-Host", "otp-authenticator.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	defer res.Body.Close()

	// Create a new file in the current directory
	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatalf("Failed to copy response body to file: %v", err)
	}

	fmt.Println("Secret key:", s.SecretKey)
	fmt.Println("QR code saved to qrcode.png")
	return s
}

func InitializeSecret() *SecretKeyStruct {
	s := &SecretKeyStruct{}
	s.RandomSecretLength()
	s.GenerateSecretKey()
	s.CreateQRCode()
	return s
}
