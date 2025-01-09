package concurency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fastest server wins if under 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("got unexpected error %q", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("return an error if server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)
		defer serverA.Close()
		defer serverB.Close()

		serverAURL := serverA.URL
		serverBURL := serverB.URL

		_, err := ConfigurableRacer(serverAURL, serverBURL, 100*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
