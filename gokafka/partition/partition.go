package partition

type PartitionInfo struct {
	ID       uint32
	Topic    string
	IsLeader bool
}

type Partitioner interface {
	Partition(key interface{}, numPartitions int) int
}
