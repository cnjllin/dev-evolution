// channel project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done //阻塞，直到前面的执行完

	pingmsg := make(chan string, 1)
	pongmsg := make(chan string, 1)

	ping(pingmsg, "hhhhhh,this is so redicuous")
	pong(pingmsg, pongmsg)

	fmt.Println(<-pongmsg)

	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(ping chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}
