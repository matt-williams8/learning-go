package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestWebsiteChecker(t *testing.T) {
	websiteUrls := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(func(websiteUrl string) bool {
		if websiteUrl == "waat://furhurterwe.geds" {
			return false
		} else {
			return true
		}
	}, websiteUrls)

	if reflect.DeepEqual(want, got) != true {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func BenchmarkWebsiteChecker(b *testing.B) {

	testWebsiteUrls := make([]string, 100)
	for i := 0; i < len(testWebsiteUrls); i++ {
		testWebsiteUrls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(func(_ string) bool {
			// slow stub to replicate network latency?
			time.Sleep(20 * time.Millisecond)
			return true
		}, testWebsiteUrls)
	}

}
