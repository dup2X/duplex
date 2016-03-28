//+build ignore

package main

import (
	"fmt"
	"net"
	"sync"
	"syscall"
)

const (
	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
	BACKLOG        = 128
	PAGE_SIZE      = 4 * 1024
)

type Message struct {
	conn int
	data []byte
}

type TCPSvr struct {
	sync.RWMutex
	listenFd int
	epFd     int
	conns    map[int]struct{}
	inbs     chan *Message
}

func (s *TCPSvr) GetConn(conn int) (ok bool) {
	s.RLock()
	defer s.RUnlock()
	_, ok = s.conns[conn]
	return
}

func (s *TCPSvr) DelConn(conn int) {
	s.Lock()
	defer s.Unlock()
	delete(s.conns, conn)
}

func (s *TCPSvr) PutConn(conn int) {
	s.Lock()
	defer s.Unlock()
	s.conns[conn] = struct{}{}
}

func (s *TCPSvr) sendLoop() {
	for {
		select {
		case msg := <-s.inbs:
			fmt.Printf("recv input %v\n", msg)
			s.writeFd(msg.conn)
		}
	}
}

func (s *TCPSvr) run() {
	var event syscall.EpollEvent
	var events [MaxEpollEvents]syscall.EpollEvent
	var err error

	s.listenFd, err = syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer syscall.Close(s.listenFd)

	if err = syscall.SetNonblock(s.listenFd, true); err != nil {
		fmt.Println("setnonblock1: ", err)
		return
	}

	addr := syscall.SockaddrInet4{Port: 10086}
	copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())

	syscall.Bind(s.listenFd, &addr)
	syscall.Listen(s.listenFd, BACKLOG)

	s.epFd, err = syscall.EpollCreate1(0)
	if err != nil {
		fmt.Println("epoll_create1: ", err)
		return
	}
	defer syscall.Close(s.epFd)

	event.Events = syscall.EPOLLIN
	event.Fd = int32(s.listenFd)
	if err = syscall.EpollCtl(s.epFd, syscall.EPOLL_CTL_ADD, s.listenFd, &event); err != nil {
		fmt.Println("epoll_ctl: ", err)
		return
	}

	for {
		nevents, e := syscall.EpollWait(s.epFd, events[:], -1)
		if e != nil {
			fmt.Println("epoll_wait: ", e)
			break
		}

		for ev := 0; ev < nevents; ev++ {
			if int(events[ev].Fd) == s.listenFd {
				connFd, _, err := syscall.Accept(s.listenFd)
				if err != nil {
					fmt.Println("accept: ", err)
					continue
				}
				fmt.Printf("accept new fd: %d\n", connFd)
				syscall.SetNonblock(connFd, true)
				event.Events = syscall.EPOLLIN | EPOLLET
				event.Fd = int32(connFd)
				if err := syscall.EpollCtl(s.epFd, syscall.EPOLL_CTL_ADD, connFd, &event); err != nil {
					fmt.Printf("epoll_ctl:%d %v\n", connFd, err)
					continue
				}
				s.PutConn(connFd)
			} else if events[ev].Events&syscall.EPOLLIN > 0 {
				fmt.Printf("event = %d\n", events[ev].Events)
				sockFd := events[ev].Fd
				if sockFd < 0 {
					continue
				}
				fmt.Sprintf("in fd=%d\n", sockFd)
				closed := s.readFd(int(sockFd))
				fmt.Printf("read over closed: %v\n", closed)
				if closed {
					s.DelConn(int(sockFd))
				}
			} else if events[ev].Events&syscall.EPOLLOUT > 0 {
				println("out")
			} else {
				println("unknown")
			}
		}
	}

}

func (s *TCPSvr) writeFd(fd int) {
	if ok := s.GetConn(fd); !ok {
		return
	}
	res := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\nHello World", 11)
	var buf = []byte(res)
	n, err := syscall.Write(fd, buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("send to %d %d bytes\n", fd, n)
}

func (s *TCPSvr) readFd(fd int) (closed bool) {
	var buf [PAGE_SIZE]byte
	var n int
	for {
		nbytes, err := syscall.Read(fd, buf[:])
		if nbytes > 0 {
			fmt.Printf("read %d >>> %s\n", fd, buf)
		}
		fmt.Printf("read %d from %d err:%v\n", nbytes, fd, err)
		if err != syscall.EAGAIN && nbytes == -1 {
			fmt.Printf("closed %d\n", fd)
			syscall.Close(fd)
			closed = true
			return
		}
		if nbytes == 0 {
			fmt.Printf("closed1 %d\n", fd)
			closed = true
			syscall.Close(fd)
			return
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if nbytes <= 0 {
			break
		} else {
			n = nbytes
		}
	}
	s.inbs <- &Message{fd, buf[:n]}
	return
}

func main() {
	s := &TCPSvr{}
	s.conns = make(map[int]struct{})
	s.inbs = make(chan *Message, 1024)
	go s.sendLoop()
	s.run()
}
