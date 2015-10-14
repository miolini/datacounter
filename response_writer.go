package datacounter

import (
	"io"
	"net/http"
	"sync/atomic"
)

type ResponseWriterCounter struct {
	http.ResponseWriter
	count  uint64
	writer io.Writer
}

func NewResponseWriterCounter(rw http.ResponseWriter) *ResponseWriterCounter {
	return &ResponseWriterCounter{
		writer: rw,
	}
}

func (counter *ResponseWriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.writer.Write(buf)
	if err != nil {
		return 0, err
	}
	atomic.AddUint64(&counter.count, uint64(n))
	return n, nil
}

func (counter *ResponseWriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}
