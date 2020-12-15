package gofocused

import (
	"crypto/sha1"
	"fmt"
	"time"
)

var logText = "I Love Your Code"

// LetsSha1Hash use sha1 hash that mean text your input cannot be back anymore
func LetsSha1Hash() {
	var sha = sha1.New()
	sha.Write([]byte(logText))

	var encrypted = sha.Sum(nil)
	var encryptedStr = fmt.Sprintf("%s", encrypted)
	fmt.Println(encryptedStr)
}

func UseSha1Salting() {
	fmt.Printf("Original: %s\n\n", logText)

	var hashed1, _ = sha1Salting(logText)
	fmt.Printf("Hashed 1: %s\n\n", hashed1)

	var hashed2, _ = sha1Salting(logText)
	fmt.Printf("Hashed 2: %s\n\n", hashed2)

	var hashed3, _ = sha1Salting(logText)
	fmt.Printf("Hashed 3: %s\n\n", hashed3)
}

func sha1Salting(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("Text: '%s', Salt: %s", text, salt)
	fmt.Println(saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))

	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
