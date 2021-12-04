package main

import (
	"fmt"
	"sync"
	"time"
)

var j int = 5

func producer(mu *sync.Mutex, c *chan int) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		fmt.Println("Producing started")
		if len(*c) == 5 {
			fmt.Println("Producing finished")
			mu.Unlock()
		} else {
			*c <- j
			j = j + 1
			fmt.Println("Producing finished")
			mu.Unlock()
		}
	}
}

func consumer(mu *sync.Mutex, c *chan int) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		fmt.Println("Consuming started")
		if len(*c) == 0 {
			fmt.Println("Consuming finished")
			mu.Unlock()
		} else {
			fmt.Println(<-(*c))
			fmt.Println("Consuming finished")
			mu.Unlock()
		}
	}
}

func main() {
	c := make(chan int, 5)
	var mu sync.Mutex
	go producer(&mu, &c)
	go consumer(&mu, &c)
	time.Sleep(10 * time.Millisecond)
}
