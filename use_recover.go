package gofocused

import (
	"fmt"
)

func UseRecover() {
	defer catch()

	var name string

	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		panic(err.Error())
		fmt.Println("End")
	}
}

func catch() {
	if rec := recover(); rec != nil {
		fmt.Println("Error occured -", rec)
	} else {
		fmt.Println("App running perfectly")
	}
}
