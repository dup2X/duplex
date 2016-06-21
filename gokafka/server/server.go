package server

import (
	"net"
	"time"
)

type Server struct {
	serveAddr string

	l *net.TCPListener
}

func (s *Server) Serve() {
	addr, _ := net.ResolveTCPAddr("tcp4", s.serveAddr)
	s.l, _ = net.ListenTCP("tcp4", addr)
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
	req, err := readRequest(c)
	if err != nil {
		return
	}
	s.input <- req
}

func (s *Server) run() {
	for {
		select {
		case req := <-s.input:
			switch req.Cmd {
			case GET:
				storage.Get(req)
			case PUT:
				storage.Put(req)
			}
		}
	}
}
