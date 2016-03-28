package direct

import (
	"net"
)

// indirect
var (
	lookupSRV = net.LookupSRV
)

func demo() {
	cName, addrs, err := lookupSRV("", "tcp", "")
}

type NullStruct struct{}
