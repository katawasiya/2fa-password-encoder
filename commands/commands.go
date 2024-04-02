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
		_, err := crypt.Encrypt(filename)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
	case 2:
		fmt.Println("Decrypt")
		fmt.Print("Enter filename: ")
		var filename string
		fmt.Scanln(&filename)
		fmt.Print("Enter decryption key: ")
		var decryptionKey string
		fmt.Scanln(&decryptionKey)
		_, err := crypt.Decrypt(decryptionKey, filename)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Invalid Command")
	}
}
