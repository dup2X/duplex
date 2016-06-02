package consumer

type Consumer struct {
	Group     string
	OffsetMgr map[string]uint64
}
