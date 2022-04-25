package racer

import (
	"fmt"
	"net/http"
	"time"
)

func UrlRacer(url1, url2 string) (fastestUrl string, err error) {
	return ConfigurableUrlRacer(url1, url2, 10*time.Second)
}

func ConfigurableUrlRacer(url1, url2 string, timeout time.Duration) (fastestUrl string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting %v for %s and %s", timeout, url1, url2)
	}
}

func ping(url string) chan struct{} {
	repsonseChannel := make(chan struct{})

	go func() {
		http.Get(url)
		close(repsonseChannel)
	}()

	return repsonseChannel
}
