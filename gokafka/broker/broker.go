package broker

type ProtocolType uint8

type Broker struct {
	ID        int
	EndPoints map[ProtocolType]EndPoint
	Rack      string
}

func NewBroker(id int, endPoints map[ProtocolType]EndPoint, rack string) *Broker {
	return &Broker{
		ID:        id,
		EndPoints: endPoints,
		Rack:      rack,
	}
}

type EndPoint struct {
	Host string
	Port int
}
