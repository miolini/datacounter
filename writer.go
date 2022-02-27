package datacounter

import (
	"io"
	"sync/atomic"
)

// WriterCounter is counter for io.Writer.
type WriterCounter struct {
	io.Writer
	count  uint64
	writes uint64
}

// NewWriterCounter function creates a new WriterCounter.
func NewWriterCounter(w io.Writer) *WriterCounter {
	return &WriterCounter{
		Writer: w,
		count:  0,
		writes: 0,
	}
}

// Write counts bytes written and increments write counter.
func (counter *WriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.Writer.Write(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	atomic.AddUint64(&counter.writes, uint64(1))

	return n, err //nolint:wrapcheck
}

// Count function returns counted bytes.
func (counter *WriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Writes function returns count of Write() calls.
func (counter *WriterCounter) Writes() uint64 {
	return atomic.LoadUint64(&counter.writes)
}
