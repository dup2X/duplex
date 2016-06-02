package topic

type Topic struct {
	ID    string
	Group GroupSt

	MaxPartitionSize uint64
	PartitionNum     int
	PartitionIndex   uint32 // Write
}

type GroupSt string
