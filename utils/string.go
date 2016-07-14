package utils

import (
	"reflect"
	"unsafe"
)

func String(data []byte) (str string) {
	if len(data) == 0 {
		return ""
	}

	ptrBytes := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	ptrString := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	ptrString.Data = ptrBytes.Data
	ptrString.Len = ptrBytes.Len
	return
}
