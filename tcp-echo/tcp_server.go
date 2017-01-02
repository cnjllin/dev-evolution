package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type Numbers struct {
	X int
	Y int
}

func main() {
	fmt.Println("tcp server is starting ...")
	tcpAddr, e := net.ResolveTCPAddr("tcp", "localhost:9090")
	if e != nil {
		fmt.Printf("create conn occur error:%v", e)
	}

	listen, e := net.ListenTCP("tcp", tcpAddr) // 一直监听
	if e != nil {
		fmt.Printf("create conn occur error:%v", e)
	}
	tcpConn, e := listen.AcceptTCP()
	if e != nil {
		fmt.Printf("create conn occur error:%v", e)
	}
	defer tcpConn.Close()

	fmt.Printf("localAddr is %v,remoteAddr is %v\n", tcpConn.LocalAddr(), tcpConn.RemoteAddr())

	for {
		buf := make([]byte, 1024, 65535)
		n, _ := tcpConn.Read(buf)
		if n == 0 {
			break
		}
		fmt.Printf("receive from client, data is :%v\n", string(buf[:n]))

		nums := Numbers{}
		e = json.Unmarshal(buf[:n], &nums)
		if e != nil {
			fmt.Printf("json Unmarshal error:%v\n", e)
			os.Exit(-1)
		}

		var ret int
		ret = nums.X + nums.Y

		fmt.Printf("x+y=%d\n", ret)
		n, e = tcpConn.Write([]byte(strconv.Itoa(int(ret))))
		if e != nil {
			fmt.Printf("write data to client error:%v\n", e)
		}
		if n == 0 {
			fmt.Println("n is 0")
		}
		fmt.Printf("send to client,data is :%d\n", ret)

		time.Sleep(3000 * time.Millisecond)
		fmt.Println()
	}
	fmt.Println("server graceful shutdown")
}
