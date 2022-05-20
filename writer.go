package datacounter

import (
	"io"
	"sync/atomic"
)

// WriterCounter is counter for io.Writer
type WriterCounter struct {
	count uint64
	io.Writer
}

// NewWriterCounter function create new WriterCounter
func NewWriterCounter(w io.Writer) *WriterCounter {
	return &WriterCounter{
		Writer: w,
	}
}

func (counter *WriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.Writer.Write(buf)

	// Write() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Writer may return negative
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
func (counter *WriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}
