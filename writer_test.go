package datacounter_test

import (
	"bytes"
	"testing"

	"golift.io/datacounter"
)

func TestWriterCounter(t *testing.T) {
	t.Parallel()

	var (
		data    = []byte("Hello, World!")
		dataLen = uint64(len(data))
	)

	buf := bytes.Buffer{}
	counter := datacounter.NewWriterCounter(&buf)
	_, _ = counter.Write(data)

	if counter.Count() != dataLen {
		t.Fatalf("count mismatch len of test data: %d != %d", counter.Count(), len(data))
	}

	if counter.Writes() != 1 {
		t.Fatalf("write mismatch len of test data: %d != 1", counter.Writes())
	}
}
