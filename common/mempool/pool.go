package mempool

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
	ch   chan *bytes.Buffer
	cap  int
}

func NewBufferPool(poolSize, cap int) *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, cap))
			},
		},
		ch:  make(chan *bytes.Buffer, poolSize),
		cap: cap,
	}
}

func (bp *BufferPool) Get() *bytes.Buffer {
	select {
	case bf := <-bp.ch:
		return bf
	default:
		return bp.pool.Get().(*bytes.Buffer)
	}
}

func (bp *BufferPool) Put(bf *bytes.Buffer) {
	bf.Reset()
	if cap(bf.Bytes()) > bp.cap {
		old := bf
		bp.pool.Put(old)
		bf = bp.pool.Get().(*bytes.Buffer)
	}
	select {
	case bp.ch <- bf:
	default:
		bp.pool.Put(bf)
	}
}
