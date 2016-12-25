package main

import (
	"fmt"
	"net"
)

func main() {
	//DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)

	addr, e := net.ResolveUDPAddr("udp", "localhost:9090")
	if e != nil {
		fmt.Printf("resolve udp addr error :%v\n", e)
	}

	conn, e := net.DialUDP("udp", nil, addr)
	if e != nil {
		fmt.Printf("create udp conn occur errors: %s\n", e)
	}
	defer conn.Close()

	conn.Write([]byte("hello udp server"))
	fmt.Println("upd client finish")

}
