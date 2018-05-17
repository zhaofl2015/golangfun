package main

import "fmt"

import "time"

//import "golang.org/x/net/context"

//import "sync"

func worker(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		// 正常工作
		case <-cannel:
			// 退出
		}
	}
}
func main() {
	cancel := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
}
