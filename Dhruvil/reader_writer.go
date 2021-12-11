package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rw sync.RWMutex
var r = make(chan int)
var w = make(chan int)
var reader, writer int

func read() {
	time.Sleep(2 * time.Millisecond)
	rw.RLock()
	reader += 1
	r <- 1
	time.Sleep(2 * time.Millisecond)
	reader -= 1
	r <- -1
	rw.RUnlock()
	wg.Done()
}

func write() {
	time.Sleep(1 * time.Millisecond)
	rw.Lock()
	writer += 1
	w <- 1
	time.Sleep(1 * time.Millisecond)
	writer -= 1
	w <- -1
	rw.Unlock()
	wg.Done()
}

func main() {
	go func() {
		for {
			select {
			case <-r:
			case <-w:
			}
			fmt.Printf("%d reader(s)||%d writer(s)\n", reader, writer)

		}
	}()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}

// Sample o/p when run on go playground
// 0 readers||1 writers
// 0 readers||0 writers
// 1 readers||0 writers
// 2 readers||0 writers
// 1 readers||0 writers
// 0 readers||1 writers
// 0 readers||1 writers
// 0 readers||0 writers
// 1 readers||0 writers
// 2 readers||0 writers
// 3 readers||0 writers
// 2 readers||0 writers
// 1 readers||0 writers
// 0 readers||1 writers
// 0 readers||1 writers
