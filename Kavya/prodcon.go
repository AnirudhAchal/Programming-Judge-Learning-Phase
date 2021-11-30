package main

import (
	"fmt"
	"sync"
	"time"
)

var ch chan int = make(chan int, 10)
var quit chan bool = make(chan bool)

func producer() {
	var mu sync.Mutex

	for i := 1; i <= 10; i++ {
		mu.Lock()

		if len(ch) == cap(ch) {
			mu.Unlock()
		} else {
			fmt.Printf("Sent %v\n", i)
			ch <- i
			mu.Unlock()
		}
		time.Sleep(5 * time.Millisecond)
	}
	quit <- true
}

func consumer() {
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		mu.Lock()

		if len(ch) == 0 {
			mu.Unlock()
		} else {
			v := <-ch
			fmt.Printf("received %v \n", v)
			mu.Unlock()
		}
		time.Sleep(time.Millisecond)
	}
}
func main() {
	go producer()
	go consumer()
	<-quit
}
