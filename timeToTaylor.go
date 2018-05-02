package main

import (
	"fmt"
	"time"
)

func consume(time, amount int64) (units, seconds int64) {
	seconds = time
	for {
		if seconds < amount {
			break
		}
		seconds -= amount
		units += 1
	}
	return
}

const day = 86400
const hour = 3600
const minute = 60

func timeToDate(sec int64) (days, hours, minutes, seconds int64) {
	seconds = sec
	days, seconds = consume(seconds, day)
	hours, seconds = consume(seconds, hour)
	minutes, seconds = consume(seconds, minute)
	return
}

func main() {
	loc, _ := time.LoadLocation("UTC")

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	concertTime, _ := time.ParseInLocation(longForm, "Jul 7, 2017 at 7:00pm (EDT)", loc)

	fmt.Printf("Taylor Swift is in Columbus, Ohio on %v, %v %vth, %v at 7pm\n", concertTime.Weekday(), concertTime.Month(), concertTime.Day(), concertTime.Year())

	if time.Now().Before(concertTime) {
		x := time.Until(concertTime)
		d, h, m, s := timeToDate(int64(x.Seconds()))
		fmt.Printf("The time until the concert is in [%d days, %d hours, %d minutes, %d seconds]\n", d, h, m, s)
	} else {
		fmt.Println("The concert was awesome!")
	}
}
