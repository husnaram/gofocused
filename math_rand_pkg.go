package gofocused

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func StringRand(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func MathRandPkg() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Random 1st:", rand.Int())
	fmt.Println("Random 2nd:", rand.Int())
	fmt.Println("Random 3rd:", rand.Int())
	fmt.Println("Random 4th:", rand.Int())
}
