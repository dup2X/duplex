package protocol

type Message struct {
	Len     uint32
	Magic   byte
	Crc     uint32
	Payload []byte
}
