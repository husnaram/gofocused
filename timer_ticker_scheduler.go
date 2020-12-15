package gofocused

import (
	"fmt"
	"time"
)

func TimeToSleep() {
	fmt.Println("START")
	time.Sleep(time.Second * 4)
	fmt.Println("AFTER 4 SECONDS")
}

func UsingScheduler() {
	for true {
		fmt.Println("Hello!!")
		time.Sleep(time.Second * 3)
	}
}

func TheNewTimer() {
	var timer = time.NewTimer(4 * time.Second)

	fmt.Println("START")
	fmt.Println(<-timer.C)
	fmt.Println("FINISH")
}

func TheAfterFunc() <-chan string {
	var ch = make(chan string)

	time.AfterFunc(4*time.Second, func() {
		ch <- "Expired"
	})

	return ch
}

func SchedulerWithTicker() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	// This goroutine run to sleeping for 10 seconds
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	// And this section running with print out of result NewTicker()
	// until goroutine previous wake up.
	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello!!", t)
		}
	}
}
