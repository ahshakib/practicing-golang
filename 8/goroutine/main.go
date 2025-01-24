// go routine cancellation example
package main

import (
	"fmt"
	"time"
)

func worker(cancel <-chan struct{}) {
	fmt.Println("Worker started")
	defer fmt.Println("Worker stopped")
	for {
		select {
		case <-cancel:
			fmt.Println("Worker cancelled")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Working...")
		}
	}
}

func main() {
	cancel := make(chan struct{})
	go worker(cancel)
	time.Sleep(5 * time.Second)
	fmt.Println("Cancelling Worker")
	close(cancel)
	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine stopped")
}