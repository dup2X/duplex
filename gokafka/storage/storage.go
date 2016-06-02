package storage

import (
	"fmt"
	"syscall"
)

func Get(dst int, src int, offset int64, count int) {
	n, err := syscall.Sendfile(dst, src, &offset, count)
	if err != nil {
		fmt.Printf("Sendfile err : %s\n", err)
		return
	}
	fmt.Printf("sent %d bytes\n", n)
	return
}

func Put() {}
