package gofocused

import (
	"errors"
	"fmt"
	"strings"
)

func CustomError() {
	var name string

	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}
	return true, nil
}
