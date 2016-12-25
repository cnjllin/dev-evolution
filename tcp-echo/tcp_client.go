package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("vim-go")

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:9090")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("occur error:%v", err)
	}
	defer conn.Close()
	conn.SetKeepAlive(true)

	conn.Write([]byte("hello tcp server"))

	fmt.Println("tcp client finish")
}
