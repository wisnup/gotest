package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return false
}

func mockWebsiteChecker(url string) bool {

	if url != "waat://furhurterwe.geds" {
		return true
	}

	return false
}

func TestWebsiteCheck(t *testing.T) {

	t.Run("test sites", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		got := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubChecker, urls)
	}
}
