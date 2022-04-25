package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUrlRacer(t *testing.T) {

	t.Run("returns the faster url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		got, err := UrlRacer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("got unexpected error %q", err)
		}

		if got != fastUrl {
			t.Errorf("got %q, wanted %q", got, fastUrl)
		}
	})

	t.Run("returns an error if neither url loads in 10 seconds", func(t *testing.T) {
		slowServer := makeDelayedServer(25 * time.Millisecond)
		defer slowServer.Close()

		_, err := ConfigurableUrlRacer(slowServer.URL, slowServer.URL, 20*time.Millisecond)

		if err == nil {
			t.Fatal("Exepcetd an error due to urls taking too long to respond, but did not get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		responseWriter.WriteHeader(http.StatusOK)
	}))
}
