package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.RWMutex

var j int = 5

func writer_action(i int, wg *sync.WaitGroup) {
	x := (rand.Intn(100) % 5) + 1
	sec := time.Duration(x) * time.Millisecond
	mu.Lock()
	fmt.Println("writing started", i)
	time.Sleep(sec)
	fmt.Println("writing finished", i)
	mu.Unlock()
	wg.Done()
}

func reader_action(i int, wg *sync.WaitGroup) {
	x := (rand.Intn(100) % 5) + 1
	sec := time.Duration(x) * time.Millisecond
	mu.RLock()
	fmt.Println("reading started", i)
	time.Sleep(sec)
	fmt.Println("reading finished", i)
	mu.RUnlock()
	wg.Done()
}

func readers(j int, wg *sync.WaitGroup) {
	var wg1 sync.WaitGroup
	for i := 0; i < j; i++ {
		x := (rand.Intn(100) % 5) + 1
		sec := time.Duration(x) * time.Millisecond
		time.Sleep(sec)
		wg1.Add(1)
		go reader_action(i, &wg1)
	}
	wg1.Wait()
	wg.Done()
}

func writers(j int, wg *sync.WaitGroup) {
	var wg1 sync.WaitGroup
	for i := 0; i < j; i++ {
		x := (rand.Intn(100) % 5) + 1
		sec := time.Duration(x) * time.Millisecond
		time.Sleep(sec)
		wg1.Add(1)
		go writer_action(i, &wg1)
	}
	wg1.Wait()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go writers(10, &wg)
	wg.Add(1)
	go readers(10, &wg)
	wg.Wait()
}
