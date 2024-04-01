package twoFactor

import (
	"fmt"
	"math/rand"
)

type TwoFactor struct {
	QRCode string
}

type TwoFactorInterface interface {
	GenerateQRCode() string
}

// this function is generating a random QR code
func (t *TwoFactor) GenerateQRCode() string {
	// generating a random QR code
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 25)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	t.QRCode = string(b)
	return t.QRCode
}

// this function is creating a new TwoFactor object
func NewTwoFactor() *TwoFactor {
	t := &TwoFactor{}
	t.GenerateQRCode()
	fmt.Println(t.QRCode)
	return t
}
