package datacounter

import (
	"io"
	"sync/atomic"
)

type ReaderCounter struct {
	io.Reader
	count  uint64
	reader io.Reader
}

func NewReaderCounter(r io.Reader) *ReaderCounter {
	return &ReaderCounter{
		reader: r,
	}
}

func (counter *ReaderCounter) Read(buf []byte) (int, error) {
	n, err := counter.reader.Read(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	return n, err
}

func (counter *ReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}
