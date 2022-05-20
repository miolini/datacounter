package datacounter

import (
	"io"
	"sync/atomic"
)

// ReaderCounter is counter for io.Reader
type ReaderCounter struct {
	count uint64
	io.Reader
}

// NewReaderCounter function for create new ReaderCounter
func NewReaderCounter(r io.Reader) *ReaderCounter {
	return &ReaderCounter{
		Reader: r,
	}
}

func (counter *ReaderCounter) Read(buf []byte) (int, error) {
	n, err := counter.Reader.Read(buf)

	// Read() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Reader may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))
	}

	return n, err
}

// Count function return counted bytes
func (counter *ReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}
