package secretKey

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"os"

	"github.com/mdp/qrterminal/v3"
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
	CreateQRCode(label, issuer string) *SecretKeyStruct
}

func (s *SecretKeyStruct) RandomSecretLength() *SecretKeyStruct {
	validLengths := []int{16, 24, 32}
	s.SecretLength = validLengths[rand.Intn(len(validLengths))]

	return s
}

func (s *SecretKeyStruct) GenerateSecretKey() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, s.SecretLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	s.SecretKey = string(b)
	return s.SecretKey
}

func InitializeSecret() *SecretKeyStruct {
	s := &SecretKeyStruct{}
	s.RandomSecretLength()
	s.GenerateSecretKey()
	s.CreateQRCode("encrypt", "katawasiya")
	return s
}

// type TwoFactorInterface interface {
// 	GetQRCode() string
// }

func (s *SecretKeyStruct) CreateQRCode(label, issuer string) (*SecretKeyStruct, error) {
	// Создаем URI для Google Authenticator
	secret := s.SecretKey
	fmt.Println("Original Secret:", secret)

	// Преобразуем секретный ключ в base32
	ConvertToBase32 := func(s string) string {
		b := []byte(s)
		base32Secret := base32.StdEncoding.EncodeToString(b)
		return base32Secret
	}

	base32Secret := ConvertToBase32(secret)
	fmt.Println("Secret in base32:", base32Secret)

	uri := fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s", label, base32Secret, issuer)

	// Выводим QR-код в терминал
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(uri, config)

	return s, nil
}
