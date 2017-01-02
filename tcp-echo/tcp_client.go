package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Numbers struct {
	X int
	Y int
}

func main() {
	fmt.Println("tcp_client is starting ...")

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:9090") // 远端地址

	conn, err := net.DialTCP("tcp", nil, tcpAddr) // 本地地址随机分配
	if err != nil {
		fmt.Printf("occur error:%v", err)
	}
	defer conn.Close()
	conn.SetKeepAlive(true)

	fmt.Printf("localAddr is %v, remoteAddr is %v\n", conn.LocalAddr(), conn.RemoteAddr())

	var count int64 = 0
	for {
		nums := Numbers{}
		fmt.Scan(&nums.X, &nums.Y)

		fmt.Printf("x value is:%d,y value is %d\n", nums.X, nums.Y)

		data, e := json.Marshal(&nums)
		if e != nil {
			fmt.Printf("json marshal error:%v\n", e)
			os.Exit(-1)
		}

		n, e := conn.Write(data)
		if e != nil {
			fmt.Printf("write data to remote error:%v\n", e)
			os.Exit(-1)
		}
		fmt.Printf("send to remote server ,data is:%v\n", nums)

		buf := make([]byte, 1024, 65535)
		n, e = conn.Read(buf)

		var res int
		e = json.Unmarshal(buf[:n], &res)
		if e != nil {
			fmt.Printf("json Unmarshal error:%v\n", e)
		}
		fmt.Printf("two numbers sum is:%d\n", res)

		count += 1
		fmt.Printf("the %d times recicle\n", count)

	}

	fmt.Printf("tcp client finish,total count is:%d\n", count)
}
