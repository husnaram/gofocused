package gofocused

import (
	"fmt"
	"io"
	"os"
)

var path = "/home/mandasia/codes/go/src/basicme.go/the_test.txt"

func ReadFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Ready to reading content file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("--> Content file has been read")
	fmt.Println(string(text))
}

func EditingContentFile() {
	// Open and give access the file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// And ready for writing
	_, err = file.WriteString("Hallo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Lets start to coding")
	if isError(err) {
		return
	}

	// Anything you write will be store
	// Sync method similar commit git
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("--> Content file has existed")
}

func CreateFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("--> File has been created", path)
}

func DeteleFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("--> File has deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
