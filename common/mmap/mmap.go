package mmap

import (
	"syscall"
)

type ProtFlags uint

const (
	PROT_NONE  ProtFlags = 0x0
	PROT_READ  ProtFlags = 0x1
	PROT_WRITE ProtFlags = 0x2
	PROT_EXEC  ProtFlags = 0x4
)

type MapFlags uint

const (
	MAP_SHARED    MapFlags = 0x1
	MAP_PRIVATE   MapFlags = 0x2
	MAP_FIXED     MapFlags = 0x10
	MAP_ANONYMOUS MapFlags = 0x20
	MAP_GROWSDOWN MapFlags = 0x100
	MAP_LOCKED    MapFlags = 0x2000
	MAP_NONBLOCK  MapFlags = 0x10000
	MAP_NORESERVE MapFlags = 0x4000
	MAP_POPULATE  MapFlags = 0x8000
)

type SyncFlags uint

const (
	MS_SYNC       SyncFlags = 0x4
	MS_ASYNC      SyncFlags = 0x1
	MS_INVALIDATE SyncFlags = 0x2
)

type AdviseFlags uint

const (
	MADV_NORMAL     AdviseFlags = 0x0
	MADV_RANDOM     AdviseFlags = 0x1
	MADV_SEQUENTIAL AdviseFlags = 0x2
	MADV_WILLNEED   AdviseFlags = 0x3
	MADV_DONTNEED   AdviseFlags = 0x4
	MADV_REMOVE     AdviseFlags = 0x9
	MADV_DONTFORK   AdviseFlags = 0xa
	MADV_DOFORK     AdviseFlags = 0xb
)

func Mmap(addr, length, prot, flags, fd uintptr, offset int64) (uintptr, error) {
	addr, _, err := syscall.Syscall6(syscall.SYS_MMAP, addr, length, prot, flags, fd, uintptr(offset))
	if err != 0 {
		return 0, syscall.Errno(err)
	}
	return addr, nil
}

func Unmap(addr, l uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_MUNMAP, addr, l, 0)
	if err != 0 {
		return syscall.Errno(err)
	}
	return nil
}

func Madvise(addr, length, flag uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_MADVISE, addr, length, flag)
	if err != 0 {
		return syscall.Errno(err)
	}
	return nil
}
