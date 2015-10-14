package datacounter

import (
	"testing"
	"bytes"
)

func TestWriterCounter(t *testing.T) {
	buf := bytes.Buffer{}
	counter := NewWriterCounter(&buf)
	counter.Write(data)
	if counter.Count() != dataLen {
		t.Fatal("count mismatch len of test data: %d != %d", counter.Count, len(data))
	}
}