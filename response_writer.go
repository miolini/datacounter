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
	count      uint64
	started    time.Time
	statusCode int
}

// NewResponseWriterCounter function create new ResponseWriterCounter
func NewResponseWriterCounter(rw http.ResponseWriter) *ResponseWriterCounter {
	return &ResponseWriterCounter{
		ResponseWriter: rw,
		started:        time.Now(),
	}
}

// Write returns underlying Write result, while counting data size
func (counter *ResponseWriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.ResponseWriter.Write(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	return n, err
}

// Header returns underlying Header result
func (counter *ResponseWriterCounter) Header() http.Header {
	return counter.ResponseWriter.Header()
}

// WriteHeader returns underlying WriteHeader, while setting Runtime header
func (counter *ResponseWriterCounter) WriteHeader(statusCode int) {
	counter.Header().Set("X-Runtime", fmt.Sprintf("%.6f", time.Since(counter.started).Seconds()))
	counter.ResponseWriter.WriteHeader(statusCode)
}

// Hijack returns underlying Hijack
func (counter *ResponseWriterCounter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return counter.ResponseWriter.(http.Hijacker).Hijack()
}

// Count function return counted bytes
func (counter *ResponseWriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Started returns started value
func (counter *ResponseWriterCounter) Started() time.Time {
	return counter.started
}

// StatusCode returns sent status code
func (counter *ResponseWriterCounter) StatusCode() int {
	return counter.statusCode
}
