package taylorTime

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

const longForm = "Jan 2, 2006 at 3:04pm (MST)"

var loc, _ = time.LoadLocation("UTC")
var concertTime, _ = time.ParseInLocation(longForm, "Jul 7, 2018 at 7:00pm (EDT)", loc)

func timeToDate(sec int64) (days, hours, minutes, seconds int64) {
	seconds = sec
	days, seconds = consume(seconds, day)
	hours, seconds = consume(seconds, hour)
	minutes, seconds = consume(seconds, minute)
	return
}

func Future() bool {
	return time.Now().Before(concertTime)
}

func Happened() bool {
	return !Future()
}

func Banner() string {
	return fmt.Sprintf("Taylor Swift is in Columbus, Ohio on %v, %v %vth, %v at 7pm", concertTime.Weekday(), concertTime.Month(), concertTime.Day(), concertTime.Year())
}

func TimeRemaining() string {
	if Future() {
		x := time.Until(concertTime)
		d, h, m, s := timeToDate(int64(x.Seconds()))
		return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", d, h, m, s)
	} else {
		return fmt.Sprint("The concert was awesome!")
	}
}
