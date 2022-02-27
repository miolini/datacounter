# Datacounter

Golang counters for readers/writers.

[![Build Status](https://travis-ci.org/golift/datacounter.svg)](https://travis-ci.org/golift/datacounter)

[GoDoc](http://godoc.org/github.com/golift/datacounter)

## Fork?

This fork adds `.Read()` and `.Write()` method counters in addition to the existing byte counters.

## Examples

### ReaderCounter

```go

import "golift.io/datacounter"

buf := bytes.Buffer{}
buf.Write(data)
counter := datacounter.NewReaderCounter(&buf)

io.Copy(ioutil.Discard, counter)
if counter.Count() != dataLen {
	t.Fatalf("count mismatch len of test data: %d != %d", counter.Count(), len(data))
}
```

### WriterCounter

```go
import "golift.io/datacounter"

buf := bytes.Buffer{}
counter := datacounter.NewWriterCounter(&buf)

counter.Write(data)
if counter.Count() != dataLen {
	t.Fatalf("count mismatch len of test data: %d != %d", counter.Count(), len(data))
}
```

### http.ResponseWriter Counter

```go
import "golift.io/datacounter"

handler := func(w http.ResponseWriter, r *http.Request) {
	w.Write(data)
}

req, err := http.NewRequest("GET", "http://example.com/foo", nil)
if err != nil {
	t.Fatal(err)
}

w := httptest.NewRecorder()
counter := datacounter.NewResponseWriterCounter(w)

handler(counter, req)
if counter.Count() != dataLen {
	t.Fatalf("count mismatch len of test data: %d != %d", counter.Count(), len(data))
}
```
