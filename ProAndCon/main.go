// ProAndCon project main.go
package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("value is:", v)
	}
}

func main() {
	fmt.Println("生产者消费者模拟")
	c := make(chan int, 1)
	go Consumer(c)
	go Producer(c)

	time.Sleep(100000000)

}
