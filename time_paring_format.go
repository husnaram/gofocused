package gofocused

import (
	"fmt"
	"time"
)

func UseTime() {
	var timeOne = time.Now()
	fmt.Printf("Time One: %v\n", timeOne)

	var timeTwo = time.Date(1995, 06, 02, 10, 30, 0, 0, time.UTC)
	fmt.Printf("Time Two: %v\n", timeTwo)

	var timeNow = time.Now()
	fmt.Println("Year:", timeNow.Year())
	fmt.Println("Month:", timeNow.Month())
	fmt.Println("Day:", timeNow.Weekday())
	fmt.Println("Location:", timeNow.Location())
}

func StringTimeToTimeParser() {
	var layoutForm, value string
	var date time.Time
	// var e

	layoutForm = "2006-01-02 15:04 PM"
	value = "2019-11-12 09:09 AM"
	date, _ = time.Parse(layoutForm, value)
	fmt.Println(value, "\t->", date.String())

	layoutForm = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutForm, value)
	fmt.Println(value, "\t\t->", date.String())

	// Time parse with predefined layout form
	date, _ = time.Parse(time.RFC822, "19 Oct 03 08:00 WIB")
	fmt.Println(date.String())
}

func TimeToString() {
	var date, err = time.Parse(time.RFC822, "02 Sep 15 08:04 WIB")
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	var dateOne = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("Date One:", dateOne)

	var dateTwo = date.Format(time.RFC3339)
	fmt.Println("Date Two:", dateTwo)
}
