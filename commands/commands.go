package commands

import (
	"bufio"
	crypt "encryptor/pckg/crypt"
	"fmt"
	"os"
	"strings"
)

func Commands() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: encryptor <file> <operation>")
		os.Exit(1)
	}

	operation := os.Args[1]
	filename := os.Args[2]

	switch operation {
	case "encrypt":
		crypt.Encrypt(filename)
	case "decrypt":
		fmt.Print("Enter decryption key: ")
		reader := bufio.NewReader(os.Stdin)
		decryptionKey, _ := reader.ReadString('\n')
		decryptionKey = strings.TrimSpace(decryptionKey)
		crypt.Decrypt(decryptionKey, filename)
	default:
		fmt.Println("Error: Invalid operation. Please enter 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}
}
