package datacounter_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"golift.io/datacounter"
)

func TestResponseWriterCounter(t *testing.T) {
	t.Parallel()

	var (
		data    = []byte("Hello, World!")
		dataLen = uint64(len(data))
	)

	handler := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(data)
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	counter := datacounter.NewResponseWriterCounter(w)
	handler(counter, req)

	if counter.Count() != dataLen {
		t.Fatalf("count mismatch len of test data: %d != %d", counter.Count(), len(data))
	}

	if counter.Writes() != 1 {
		t.Fatalf("write mismatch len of test data: %d != 1", counter.Writes())
	}
}
