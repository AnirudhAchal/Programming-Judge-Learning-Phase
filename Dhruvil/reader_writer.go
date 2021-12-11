package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var rw sync.RWMutex
var data = 10

func read() {
	for true {
		rw.RLock()
		fmt.Printf("read data: %d\n", data)
		rw.RUnlock()
	}

	wg.Done()
}

func write() {
	for true {
		rw.Lock()
		data = rand.Int() % 100
		fmt.Printf("write data: %d\n", data)
		rw.Unlock()
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	go read()
	go write()
	wg.Wait()
}
