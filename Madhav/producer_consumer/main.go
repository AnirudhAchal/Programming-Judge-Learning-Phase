package main

import (
	"fmt"
	"sync"
	"time"
)

type a struct {
	arr [5]int
	mu  sync.Mutex
	i   int
}

var j int = 6

func producer(b *a) {
	for true {
		b.mu.Lock()
		(b.i) = ((b.i) + 1) % 5
		b.arr[b.i] = j
		b.mu.Unlock()
	}
}

func consumer(b *a) {
	for true {
		b.mu.Lock()
		fmt.Println(b.arr[b.i])
		(b.i) = ((b.i) + 1) % 5
		b.mu.Unlock()
	}
}

func main() {
	var b a
	b.i = 0
	go producer(&b)
	go consumer(&b)
	time.Sleep(1 * time.Millisecond)
}
