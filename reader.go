package datacounter

import (
	"io"
	"sync/atomic"
)

// ReaderCounter is counter for io.Reader.
type ReaderCounter struct {
	count uint64
	reads uint64
	io.Reader
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
	atomic.AddUint64(&counter.reads, 1)

	// Read() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Reader may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	n, err := counter.Reader.Read(buf)
	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))
	}

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
