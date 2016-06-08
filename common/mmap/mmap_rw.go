package mmap

import (
	"errors"
	"os"
	"reflect"
	"sync"
	"unsafe"
)

var errInvalidSize = errors.New("Invalid size, must be 1024*N")

type MmapFileWriter struct {
	f    *os.File
	size int64
	addr uintptr

	mu       *sync.Mutex
	buf      []byte
	readPos  int
	writePos int
}

func NewMmapFileWriter(name string, size int64) (*MmapFileWriter, error) {
	if size < 1024 || size%1024 != 0 {
		return nil, errInvalidSize
	}
	var fw *MmapFileWriter
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	defer func() {
		if fw == nil {
			f.Close()
		}
	}()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if fi.Size() < size {
		kb := make([]byte, 1024)
		for i := int64(0); i < size/1024; i++ {
			f.Write(kb)
		}
		f.Sync()
	}

	addr, err := Mmap(uintptr(0), uintptr(size), uintptr(PROT_WRITE|PROT_READ), uintptr(MAP_SHARED), f.Fd(), int64(0))
	if err != nil {
		return nil, err
	}
	defer func() {
		if fw == nil {
			Unmap(addr, uintptr(size))
		}
	}()

	err = Madvise(addr, uintptr(size), uintptr(MADV_SEQUENTIAL))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size, size)
	dh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	dh.Len = int(size)
	dh.Cap = int(size)
	dh.Data = addr

	fw = &MmapFileWriter{
		f:    f,
		buf:  buf,
		size: size,
		addr: addr,
		mu:   new(sync.Mutex),
	}
	return fw, nil
}

func (fw *MmapFileWriter) Write(data []byte) {
	fw.mu.Lock()
	size := len(data)
	copy(fw.buf[fw.writePos:], data)
	fw.writePos += size
	fw.mu.Unlock()
}

func (fw *MmapFileWriter) Close() {
	fw.f.Close()
	Unmap(fw.addr, uintptr(fw.size))
}
