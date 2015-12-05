/**
* Server
 */

package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct { //操作数
	A, B int
}

type Quotient struct { //商
	Quo, Rem int
}

type RPC int

//定义乘法函数
func (t *RPC) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

//服务器端的除法和取模运算
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
	rpc.Register(myrpc) //注册服务
	rpc.HandleHTTP()    //通过http的rpc

	err := http.ListenAndServe(":1234", nil) //监听端口
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("正在监听1234端口")
}
