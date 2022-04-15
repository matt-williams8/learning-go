package countdown

import (
	"fmt"
	"io"
	"time"
)

type ConfigurableSleeper struct {
	Duration      time.Duration
	SleepDelegate func(time.Duration)
}

func (configurableSleeper *ConfigurableSleeper) Sleep() {
	configurableSleeper.SleepDelegate(configurableSleeper.Duration)
}

type Sleeper interface {
	Sleep()
}

const countdownStart = 3
const finalWord = "Go!"

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}
