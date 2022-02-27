package datacounter_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"golift.io/datacounter"
)

func TestReaderCounter(t *testing.T) {
	t.Parallel()

	var (
		data    = []byte("Hello, World!")
		dataLen = uint64(len(data))
	)

	buf := bytes.Buffer{}
	_, _ = buf.Write(data)
	counter := datacounter.NewReaderCounter(&buf)
	_, _ = io.Copy(ioutil.Discard, counter)

	if counter.Count() != dataLen {
		t.Fatalf("count mismatch len of test data: %d != %d", counter.Count(), len(data))
	}

	if counter.Reads() != 2 {
		t.Fatalf("reads mismatch len of test data: %d != 1", counter.Reads())
	}
}
