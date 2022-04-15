package main

import (
	"mocking/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 10 * time.Second, SleepDelegate: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
