package gateway

type WorkModeType uint8

const (
	UNKNOWN = iota
	WORK_HTTP
	WORK_HTTP2
	WORK_TCP
	WORK_WEBSOCKET
	WORK_UDP
)

type Handler interface {
}

type Gateway struct {
	ID    uint64
	Token uint32

	Mode    WorkModeType
	handler Handler
}
