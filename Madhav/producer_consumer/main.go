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
		if len(*c) == 5 {
			mu.Unlock()
		} else {
			*c <- j
			j = j + 1
			mu.Unlock()
		}
	}
}

func consumer(mu *sync.Mutex, c *chan int) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		if len(*c) == 0 {
			mu.Unlock()
		} else {
			fmt.Println(<-(*c))
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
