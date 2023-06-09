package pool

import (
	"arena"
	"bufio"
	"io"
	"sync"
)

type WriterPool interface {
	GetBufferWriter(w io.Writer) *bufio.Writer
	PutBufferWriter(bw *bufio.Writer)
	Destory()
}

type writerPool struct {
	*arena.Arena
	*sync.Pool
}

func (rp *writerPool) GetBufferWriter(r io.Writer) *bufio.Writer {
	br := rp.Pool.Get().(*bufio.Writer)
	br.Reset(r)
	return br
}

func (rp *writerPool) PutBufferWriter(br *bufio.Writer) {
	br.Reset(nil)
	rp.Pool.Put(br)
}

func (rp *writerPool) Destory() {
	rp.Free()
}
func NewWriterPool(bufSize int) WriterPool {
	a := arena.NewArena()
	p := arena.New[sync.Pool](a)
	p.New = func() any {
		return bufio.NewWriterSize(defaultReadWriter, bufSize)
	}
	return &writerPool{
		Arena: a,
		Pool:  p,
	}
}
