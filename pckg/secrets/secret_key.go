package secretKey

import (
	"fmt"
	"math/rand"
)

type SecretKeyStruct struct {
	SecretKey    string
	SecretLength int
}

type SecretKeyInterface interface {
	RandomSecretLength() *SecretKeyStruct
	GenerateSecretKey() *SecretKeyStruct
	InitializeSecret() *SecretKeyStruct
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
	fmt.Print("Secret Key: ", s.SecretKey, "\n")
	return s
}
