package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.baiyuxiong.com/",
		"http://www.qiniu.com/",
		"http://www.hao123.com/",
		"http://www.163.com/",
		"http://www.sina.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1) // add 必须放在主程序中
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done() // done 方法在其他线程中调用
			// Fetch the URL.
			http.Get(url)
			fmt.Println(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait() //等待所有线程执行完毕
	fmt.Println("over")
}
