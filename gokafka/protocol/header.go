package protocol

import (
	"bufio"
	"io"
	"net"
)

const (
	_HEADER_LEN = 0x8

	_V = 0x1
)

type PacketType uint8

const (
	UNKNOWN PacketType = iota
	WR_REQ
	WR_RESP
	RD_REQ
	RD_RESP
)

type Header struct {
	Ver     uint8
	Magic   uint16
	Type    PacketType
	BodyLen uint32
}

type Packet struct {
	Header
	Payload []byte
}

func ParsePacket(conn *net.TCPConn) {
	// TODO use mempool
	data := make([]byte, _HEADER_LEN)
	reader := bufio.NewReader(conn)
	n, err := io.ReadFull(reader, data[:1])
	if n != 1 || err != nil {
		return
	}
	if data[0] != _V {
		// TODO
		return
	}
	n, err = io.ReadFull(reader, data[1:3])
	if n != 2 || err != nil {
		return
	}
	// TODO check Magic

}
