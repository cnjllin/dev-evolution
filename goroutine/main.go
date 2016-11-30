// goroutine project main.go
package main

import (
	"fmt"
	"net/http"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello world")
}
func main() {

	f("direct method call")

	go f("go method call")

	go func(msg string) {
		fmt.Println(msg)
	}("no name method")

	//var input string ///阻塞在这里，等待go 协程执行完
	//fmt.Scanln(&input)
   http.Handle("/v1/status/activate", http.HandlerFunc(hello))
   fmt.Println(http.ListenAndServe("127.0.0.1:8089", nil))
   fmt.Println("done")


}
