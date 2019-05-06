package httprest

import (
	"net/http"
	"testing"
	"time"
)

func init() {
	go RunServer()
	time.Sleep(time.Second)
}

func BenchmarkREST(b *testing.B) {
	client := &http.Client{}

	for i := 0; i < b.N; i++ {
		_, err := client.Get("http://0.0.0.0:8001/query?name=JayChung")
		if err != nil {
			b.Fatalf("http request failed: %v", err)
		}
	}
}
