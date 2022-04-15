package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleepOperation = "sleep"
const writeOperation = "write"

type CountdownSpy struct {
	Calls []string
}

func (spy *CountdownSpy) Sleep() {
	spy.Calls = append(spy.Calls, sleepOperation)
}

func (spy *CountdownSpy) Write(bytes []byte) (n int, err error) {
	spy.Calls = append(spy.Calls, writeOperation)
	return
}

type SleepSpy struct {
	durationSlept time.Duration
}

func (sleepSpy *SleepSpy) Sleep(duration time.Duration) {
	sleepSpy.durationSlept = duration
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3, 2 , 1, Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownSpy{})

		want := `3
2
1
Go!`
		got := buffer.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("sleeps before every print", func(t *testing.T) {
		spy := &CountdownSpy{}

		Countdown(spy, spy)

		want := []string{
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
		}

		got := spy.Calls

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted calls %v, got %v", want, got)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	sleepSpy := &SleepSpy{}
	configurableSleeper := &ConfigurableSleeper{sleepTime, sleepSpy.Sleep}
	configurableSleeper.Sleep()

	got := sleepSpy.durationSlept

	if got != sleepTime {
		t.Errorf("slept for %s, wanted %s", got, sleepTime)
	}
}
