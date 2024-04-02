package commands

import (
	crypt "2fa-password-encoder/pckg/crypt"
	"fmt"
	"os"
)

func Commands() {
	fmt.Println("Commands")
	fmt.Println("1. Encrypt file")
	fmt.Println("2. Decrypt file")
	fmt.Println("3. Exit")
	fmt.Print("Enter Command: ")
	var command int
	fmt.Scanln(&command)
	switch command {
	case 1:
		fmt.Print("Enter filename: ")
		var filename string
		fmt.Scanln(&filename)
		crypt.Encrypt(filename)
	case 2:
		fmt.Println("Decrypt")
		fmt.Print("Enter decryption key: ")
		var decryptionKey string
		fmt.Scanln(&decryptionKey)
		fmt.Print("Enter filename: ")
		var filename string
		fmt.Scanln(&filename)
		crypt.Decrypt(decryptionKey, filename)
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Invalid Command")
	}
}
