# datacounter
Golang counters for readers/writers

## Examples 
### ReaderCounter
```go
	buf := bytes.Buffer{}
	buf.Write(data)
	counter := datacounter.NewReaderCounter(&buf)
	io.Copy(ioutil.Discard, counter)
	if counter.Count() != dataLen {
		t.Fatal("count mismatch len of test data: %d != %d", counter.Count, len(data))
	}`
```
### WriterCounter
```go
	buf := bytes.Buffer{}
	counter := datacounter.NewWriterCounter(&buf)
	counter.Write(data)
	if counter.Count() != dataLen {
		t.Fatal("count mismatch len of test data: %d != %d", counter.Count, len(data))
	}
```
### http.ResponseWriter Counter
```go
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
		t.Fatal("count mismatch len of test data: %d != %d", counter.Count, len(data))
	}
```
