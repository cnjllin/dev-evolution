package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("udp server is staring")

	//  func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
	addr, e := net.ResolveUDPAddr("udp", "localhost:9090")
	if e != nil {
		fmt.Printf("resolve udp addr error :%s\n", e)
	}

	conn, e := net.ListenUDP("udp", addr)
	if e != nil {
		fmt.Printf("create udp conn error:%s\n", e)
	}
	defer conn.Close()

	for {
		//		Read(b []byte) (int, error)
		buf := make([]byte, 1024, 1024)
		n, remoteAddr, e := conn.ReadFromUDP(buf)
		if e != nil {
			fmt.Printf("read data error:%v,remote addr is %v\n", e, remoteAddr)
		}
		fmt.Printf("remote addr is %v, data len is :%d, data is :%s", remoteAddr, n, string(buf[:]))
	}

	fmt.Println("udp server finish")
}
