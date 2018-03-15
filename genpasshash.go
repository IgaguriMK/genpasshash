package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/amoghe/go-crypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var saltStr string
	flag.StringVar(&saltStr, "salt", "", "Salt string")

	flag.Parse()

	encoding := base64.StdEncoding

	if saltStr == "" {
		salt := make([]byte, 12)
		_, err := rand.Read(salt)
		if err != nil {
			log.Fatal("Salt generation error:", err)
		}
		saltStr = encoding.EncodeToString(salt)
	}

	fmt.Fprint(os.Stderr, "Enter password:")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Read password failed:", err)
	}
	password := string(bytePassword)

	salt := fmt.Sprintf("$6$%s$", saltStr)
	hashStr, err := crypt.Crypt(password, salt)
	if err != nil {
		log.Fatal("Failed generating hash:", err)
	}

	fmt.Print(hashStr)
}
