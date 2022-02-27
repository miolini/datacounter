package datacounter

import (
	"io"
	"sync/atomic"
)

// ReaderCounter is counter for io.Reader.
type ReaderCounter struct {
	io.Reader
	count uint64
	reads uint64
}

// NewReaderCounter function for create new ReaderCounter.
func NewReaderCounter(r io.Reader) *ReaderCounter {
	return &ReaderCounter{
		Reader: r,
		count:  0,
		reads:  0,
	}
}

// Reads counts bytes read and increments read counter.
func (counter *ReaderCounter) Read(buf []byte) (int, error) {
	n, err := counter.Reader.Read(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	atomic.AddUint64(&counter.reads, 1)

	return n, err //nolint:wrapcheck
}

// Count function returns count of bytes read.
func (counter *ReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Reads function returns count of calls to Read() method.
func (counter *ReaderCounter) Reads() uint64 {
	return atomic.LoadUint64(&counter.reads)
}
