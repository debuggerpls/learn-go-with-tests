package concurency

import (
	"net/http"
	"time"
)

const (
	ErrTimedOut         = RacerErr("timed out")
	RacerDefaultTimeout = 10 * time.Second
)

type RacerErr string

func (r RacerErr) Error() string {
	return string(r)
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, RacerDefaultTimeout)
}

func ConfigurableRacer(a, b string, d time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(d):
		return "", ErrTimedOut
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
