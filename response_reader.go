package datacounter

import (
	"io"
	"sync/atomic"
)

// ResponseBodyCounter is counter for ReadCloser primarily targeted at wrapping http.Request Body
type ResponseBodyCounter struct {
	io.ReadCloser
	count uint64
}

// NewResponseBodyCounter function for create new ResponseBodyCounter
func NewResponseBodyCounter(r io.ReadCloser) *ResponseBodyCounter {
	return &ResponseBodyCounter{
		ReadCloser: r,
	}
}

// Read invokes the underlying Reader
func (counter *ResponseBodyCounter) Read(buf []byte) (int, error) {
	n, err := counter.ReadCloser.Read(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	return n, err
}

// Close invokes the underlying Closer
func (counter *ResponseBodyCounter) Close() error {
	return counter.ReadCloser.Close()
}

// Count function return counted bytes
func (counter *ResponseBodyCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}
