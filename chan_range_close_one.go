package gofocused

import (
	"fmt"
	"runtime"
)

func ChanRangeCloseOne() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}

func sendMessage(ch chan<- string) {
	for i := 3; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
