package mocking

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyPrintSleeper struct {
	Calls []string
}

func (s *SpyPrintSleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyPrintSleeper) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyPrintSleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spyPrintSleeper := &SpyPrintSleeper{}

		Countdown(spyPrintSleeper, spyPrintSleeper)

		calls := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(spyPrintSleeper.Calls, calls) {
			t.Errorf("wanted calls %v got %v", calls, spyPrintSleeper.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
