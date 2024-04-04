package secretKey

import (
	"fmt"
	"io"
	"log"
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

// This function generates a secret key
func (s *SecretKeyStruct) GenerateSecretKey() *SecretKeyStruct {
	apiKey := os.Getenv("RAPIDAPI_KEY")
	url := "https://otp-authenticator.p.rapidapi.com/new_v2/"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "otp-authenticator.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	s.SecretKey = string(body)
	return s
}

// This function creates a QR code
func (s *SecretKeyStruct) CreateQRCode() *SecretKeyStruct {
	url := "https://otp-authenticator.p.rapidapi.com/qr2/?data=otpauth%3A%2F%2Ftotp%2FHomeCorp%3AUser1%3Fsecret%3D" + s.SecretKey + "%26issuer%3DHomeCorp&size=5&level=M"

	payload := strings.NewReader("{}")
	apiKey := os.Getenv("RAPIDAPI_KEY")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", apiKey)
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
	s.GenerateSecretKey()
	s.CreateQRCode()
	return s
}
