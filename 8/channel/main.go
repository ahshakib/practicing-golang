package main

import (
	"fmt"
)
/* Unbuffered channel example
func greet(done chan string) {
	fmt.Println("Welcome to my app")
	done <- "Shakib"
	fmt.Println("Done...")
}

func bye(ch chan string) {
	fmt.Println("Nice to meet you.")
	time.Sleep(2 * time.Second)
	fmt.Println("Goodbye,", <-ch)
}

func main() {
	done := make(chan string)
	go greet(done)
	go bye(done)
	fmt.Scanln()
}*/

// buffered channel example
func main() {
	ch := make(chan string, 5)
	go ping(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func ping(ch chan string) {
	links := []string{"http://www.google.com", "http://www.facebook.com", "http://www.youtube.com"}
	for _, link := range links {
		ch <- link
	}
}