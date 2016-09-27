// goroutine project main.go
package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct method call")

	go f("go method call")

	go func(msg string) {
		fmt.Println(msg)
	}("no name method")

	var input string ///阻塞在这里，等待go 协程执行完
	fmt.Scanln(&input)
	fmt.Println("done")

}
