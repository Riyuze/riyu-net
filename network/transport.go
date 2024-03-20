package network

type NetAddr string

type RPC struct {
	from    string
	payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
