package gofocused

import "fmt"

// UsePanic show you how `panic` keyword will not run next line
func UsePanic() {
	var name string

	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		panic(err.Error())
		fmt.Println("End")
	}
}
