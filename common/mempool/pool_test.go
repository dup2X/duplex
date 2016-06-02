package mempool

import (
	"testing"
)

func TestBufPool(t *testing.T) {
	p := NewBufferPool(1024, 4096)
	bf := p.Get()
	if cap(bf.Bytes()) != 4096 {
		t.Log(cap(bf.Bytes()))
		t.FailNow()
	}
	p.Put(bf)
	if cap(p.ch) != 1024 {
		println(cap(p.ch))
		t.FailNow()
	}
}
