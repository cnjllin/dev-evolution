/**
* TCP  Server
 */

package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type RPC int

func (t *RPC) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *RPC) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	myrpc := new(RPC)
	rpc.Register(myrpc)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("错误： ", err.Error())
		os.Exit(1)
	}
}
