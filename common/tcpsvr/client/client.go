package main

import (
	"fmt"
	"net"
	"time"
)

var HELLO = []byte("Hello")

func do() {
	laddr, _ := net.ResolveTCPAddr("tcp4", ":0")
	raddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:10086")
	c, err := net.DialTCP("tcp4", laddr, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.Write(HELLO)
	if err != nil {
		fmt.Printf("1 %v\n", err)
		return
	}
	c.SetReadDeadline(time.Now().Add(time.Second * 2))
	buf := make([]byte, 100)
	n, err := c.Read(buf)
	if err != nil {
		fmt.Printf("%d %v\n", n, err)
		return
	}
	fmt.Println(string(buf))
	select {}
}

func main() {
	do()
	select {}
}
