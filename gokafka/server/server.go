package server

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/dup2X/duplex/gokafka/storage"
)

type Server struct {
	serveAddr string
	input     chan *Request

	l *net.TCPListener
}

func New() *Server {
	return &Server{
		serveAddr: ":10086",
		input:     make(chan *Request, 1024),
	}
}

func (s *Server) Serve() {
	addr, _ := net.ResolveTCPAddr("tcp4", s.serveAddr)
	s.l, _ = net.ListenTCP("tcp4", addr)
	go s.run()
	for i := 0; i < 4; i++ {
		go s.accept()
	}
}

func (s *Server) accept() {
	for {
		tconn, err := s.l.AcceptTCP()
		if err != nil {
			continue
		}
		tconn.SetKeepAlive(true)
		tconn.SetKeepAlivePeriod(time.Second * 300)

		go s.process(tconn)
	}
}

func (s *Server) process(c *net.TCPConn) {
	req := readRequest(c)
	s.input <- req
}

func (s *Server) run() {
	for {
		select {
		case req := <-s.input:
			fmt.Printf("get req\n")
			dst, _ := req.c.File()
			fd, _ := os.Open("/tmp/kafka_test")
			storage.Get(int(dst.Fd()), int(fd.Fd()), 0, 16)
		}
	}
}

func readRequest(c *net.TCPConn) *Request {
	return &Request{c}
}

type Request struct {
	c *net.TCPConn
}
