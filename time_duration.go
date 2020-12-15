package gofocused

import (
	"fmt"
	"time"
)

func TimeDurationPract() {
	start := time.Now()

	time.Sleep(10 * time.Second)

	duration := time.Since(start)

	fmt.Println("Time elapsed in seconds:", duration.Seconds())
	fmt.Println("Time elapsed in minutes:", duration.Minutes())
	fmt.Println("Time elapsed in houres:", duration.Hours())
}

func CalcTwoTimeObj() {
	t1 := time.Now()
	time.Sleep(20 * time.Second)
	t2 := time.Now()

	duration := t2.Sub(t1)
	fmt.Println("Time elapsed in seconds:", duration.Seconds())
	fmt.Println("Time elapsed in minutes:", duration.Minutes())
	fmt.Println("Time elapsed in houres:", duration.Hours())
}
