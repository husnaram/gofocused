package gofocused

import (
	"fmt"
	"os"
	"time"
)

func QAndA() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("What is 725/25? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("The answer is Right!")
	} else {
		fmt.Println("The answer is Wrong!")
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\nTime Out! No answer more than", timeout, "seconds")
	os.Exit(0)
}
