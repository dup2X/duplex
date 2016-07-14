package utils

import (
	"unsafe"
)

func IsBigEndian() bool {
	var i int = 0x1
	ptr := (*[int(unsafe.Sizeof(0))]byte)(unsafe.Pointer(&i))
	if ptr[0] == 0 {
		return true
	}
	return false
}
