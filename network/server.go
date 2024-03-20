package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	serverOpts ServerOpts
	rpcCh      chan RPC
	quitCh     chan struct{}
}

func NewServer(serverOpts ServerOpts) *Server {
	return &Server{
		serverOpts: serverOpts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do something")
		}
	}

	fmt.Println("Server shutting down...")
}

func (s *Server) initTransports() {
	for _, tr := range s.serverOpts.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
