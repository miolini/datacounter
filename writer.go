package datacounter

import (
	"io"
	"sync/atomic"
)

type WriterCounter struct {
	io.Writer
	count uint64
	writer io.Writer
}

func NewWriterCounter(w io.Writer) *WriterCounter {
	return &WriterCounter{
		writer: w,
	}
}

func (counter *WriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.writer.Write(buf)
	if err != nil {
		return 0, err
	}
	atomic.AddUint64(&counter.count, uint64(n))
	return n, nil
}

func (counter *WriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}