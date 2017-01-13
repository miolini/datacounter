package datacounter

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

// ResponseWriterCounter is counter for http.ResponseWriter
type ResponseWriterCounter struct {
	http.ResponseWriter
	count   uint64
	started time.Time
}

// NewResponseWriterCounter function create new ResponseWriterCounter
func NewResponseWriterCounter(rw http.ResponseWriter) *ResponseWriterCounter {
	return &ResponseWriterCounter{
		ResponseWriter: rw,
		started:        time.Now(),
	}
}

func (counter *ResponseWriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.ResponseWriter.Write(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	return n, err
}

func (counter *ResponseWriterCounter) Header() http.Header {
	return counter.ResponseWriter.Header()
}

func (counter *ResponseWriterCounter) WriteHeader(statusCode int) {
	counter.Header().Set("X-Runtime", fmt.Sprintf("%.6f", time.Since(counter.started).Seconds()))
	counter.ResponseWriter.WriteHeader(statusCode)
}

func (counter *ResponseWriterCounter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return counter.ResponseWriter.(http.Hijacker).Hijack()
}

// Count function return counted bytes
func (counter *ResponseWriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

func (counter *ResponseWriterCounter) Started() time.Time {
	return counter.started
}
