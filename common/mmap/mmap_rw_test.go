package mmap

import (
	"testing"
)

func TestMmapFileWriter(t *testing.T) {
	fw, err := NewMmapFileWriter("/tmp/mmap_test1", 1024*1024)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fw.Write([]byte("Hello"))
	fw.Close()
}
