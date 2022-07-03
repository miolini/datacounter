package datacounter

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

// ResponseWriterCounter is counter for http.ResponseWriter.
type ResponseWriterCounter struct {
	http.ResponseWriter
	count      uint64
	writes     uint64
	started    time.Time
	statusCode int
}

// NewResponseWriterCounter function create new ResponseWriterCounter.
func NewResponseWriterCounter(rw http.ResponseWriter) *ResponseWriterCounter {
	return &ResponseWriterCounter{
		ResponseWriter: rw,
		started:        time.Now(),
		statusCode:     0,
		count:          0,
		writes:         0,
	}
}

// Write returns underlying Write result, while counting data size.
func (counter *ResponseWriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.ResponseWriter.Write(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	atomic.AddUint64(&counter.writes, 1)

	return n, err //nolint:wrapcheck
}

// Header returns underlying Header result.
func (counter *ResponseWriterCounter) Header() http.Header {
	return counter.ResponseWriter.Header()
}

// WriteHeader returns underlying WriteHeader, while setting Runtime header.
func (counter *ResponseWriterCounter) WriteHeader(statusCode int) {
	counter.statusCode = statusCode
	counter.Header().Set("X-Runtime", fmt.Sprintf("%.6f", time.Since(counter.started).Seconds()))
	counter.ResponseWriter.WriteHeader(statusCode)
}

// Hijack returns underlying Hijack.
func (counter *ResponseWriterCounter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijack, exists := counter.ResponseWriter.(http.Hijacker)
	if !exists {
		return nil, nil, fmt.Errorf("writer does not support hijacking: %w", http.ErrHijacked)
	}

	return hijack.Hijack() //nolint:wrapcheck
}

// Count function return counted bytes.
func (counter *ResponseWriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Writes function returns count of Write() calls.
func (counter *ResponseWriterCounter) Writes() uint64 {
	return atomic.LoadUint64(&counter.writes)
}

// Started returns started value.
func (counter *ResponseWriterCounter) Started() time.Time {
	return counter.started
}

// StatusCode returns sent status code.
func (counter *ResponseWriterCounter) StatusCode() int {
	return counter.statusCode
}
