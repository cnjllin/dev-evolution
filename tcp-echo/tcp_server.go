package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, e := net.ResolveTCPAddr("tcp", "localhost:9090")
	if e != nil {
		fmt.Printf("create conn occur error:%v", e)
	}

	listen, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		fmt.Printf("create conn occur error:%v", e)
	}

	for {
		tcpConn, e := listen.AcceptTCP()
		if e != nil {
			fmt.Printf("create conn occur error:%v", e)
		}
		defer tcpConn.Close()

		buf := make([]byte, 1024, 1024)

		n, _ := tcpConn.Read(buf)
		fmt.Printf("data len is %d, data is :%s\n", n, string(buf[:]))

	}
}
