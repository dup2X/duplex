package producer

import (
	"github.com/dup2X/duplex/gokafka/partition"
)

type Producer struct {
	Hash partition.Partitioner
}
