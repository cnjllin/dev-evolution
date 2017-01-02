package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/astaxie/beego/logs"
)

var (
	localAddr = "localhost"
	port      = "7000"
)

func main() {
	logs.SetLogger("console")
	fmt.Println("server is starting...")
	Service()
}

func Service() {

	l, err := net.Listen("tcp", ":7000")
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			logs.Error("accept error:%v\n", err)
			return
		}
		defer conn.Close()

		fmt.Printf("localAddr is %s,remoteAddr is %s\n", conn.LocalAddr(), conn.RemoteAddr())

		go func() {
			buf := make([]byte, 65535)
			n, e := conn.Read(buf)
			if e != nil {
				logs.Error(e)
				isErr("read data error", e)
				return
			}
			fmt.Println(string(buf[:n]))
			logs.Info("receive %d bytes\n", n)
		}()

		go func() {
			n, err := conn.Write([]byte("echo client"))
			if err != nil {
				isErr("write data error", err)
				return
			}
			logs.Info("write %d bytes data to client\n", n)
		}()

	}
	logs.Info("server close")

}

func isErr(msg string, e error) {
	logs.Info(msg)
	if e != io.EOF {
		logs.Warn(msg, e)
	}
}
