/**
* 客户端调用服务器端提供的乘和取模函数
 */

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct { //操作数
	A, B int
}

type Quotient struct { //商
	Quo, Rem int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234") //默认在本地
	if err != nil {
		log.Fatal("错误:", err)
	}
	// 同步调用
	args := Args{17, 8} //两个操作数
	var reply int
	err = client.Call("RPC.Multiply", args, &reply) //通过命名空间RPC.Multiply来调用
	if err != nil {
		log.Fatal("错误:", err)
	}
	fmt.Printf("结果: %d*%d=%d\n", args.A, args.B, reply) //输出结果

	var quot Quotient
	err = client.Call("RPC.Divide", args, &quot)
	if err != nil {
		log.Fatal("错误:", err)
	}
	fmt.Printf("结果: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem) //输出结果

}
