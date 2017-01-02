package main

import (
	"fmt"
	"net"
	"os"

	"github.com/astaxie/beego/logs"
)

var (
	//log        = logs.GetLogger()
	localAddr = "127.0.0.1"
	localPort = 0 // 其实本地端口无需指定，client端的本地，只需要连接到远程的7000端口就好了
	//remoteAddr = "192.168.1.102"
	remoteAddr = "192.168.1.104"
	remotePort = "7000"
)

//proxy连接到server
func ConnectServer() (conn *net.TCPConn, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", remoteAddr+":"+remotePort)
	if err != nil {
		logs.Warn("resolve address error:%v", err)
		os.Exit(-1)
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr) //这个连接用来和远端连接，然后有数据发送过来，就发送
	if err != nil {
		logs.Warn("connect to remote server error:%v", err)
		os.Exit(-1)
	}
	conn.SetKeepAlive(true)
	return
}

//在本地new一个conn，用来和远端的某个socket连接,远端的socket是随意指定的，比如6000
//比如传来一个 是  ssh的连接，那么本地就要创建一个socket，端口由操作系统指定，然后由
//某个方法把这client的socket 和 server 的6000 端口的 socket  绑定起来，这样两个socket就
//可以互相通信了。
//func NewClientProxy(connType string) (conn *net.TCPConn, err error) {
//	listen, e := net.ListenTCP("tcp", nil)
//	if e != nil {
//		logs.Debug("client listen available port error:%v\n", e)
//		return
//	}
//	return
//}

func main() {
	conn, _ := ConnectServer()
	defer conn.Close()
	logs.Info("client is starting listen and accept")
	logs.Info("localAddr is :%s,remoteAddr is :%s\n", conn.LocalAddr(), conn.RemoteAddr())

	go func() {
		for {
			buf := make([]byte, 65535)
			n, e := conn.Read(buf)
			if e != nil {
				logs.Error("read data error:%v", e)
				fmt.Println(e)
				return
			}
			fmt.Println(string(buf[:]))

			fmt.Printf("read from remote %d bytes", n)

			buf = make([]byte, 65535)
		}
	}()

	go func() {
		for {
			n, e := conn.Write([]byte("hello server\n"))
			if e != nil {
				fmt.Printf("data write %d bytes,error is:%v \n", n, e)
			}
		}
	}()

	logs.Info("clientproxy graceful shutdown")
	return

}
