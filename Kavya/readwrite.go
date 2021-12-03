package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0

var writers int = 0
var readers int = 0
var mu sync.RWMutex

func writer(i int) {

	mu.Lock()
	writers += 1

	x = i
	time.Sleep(2 * time.Millisecond)

	writers -= 1
	mu.Unlock()
}

func reader() {

	mu.RLock()
	readers += 1

	time.Sleep(2 * time.Millisecond)

	readers -= 1
	mu.RUnlock()
}
func main() {
	var wg sync.WaitGroup

	go func() {
		for {
			fmt.Printf("%v reader(s) \n%v writer(s), x = %v\n\n", readers, writers, x)
			time.Sleep(5 * time.Millisecond)
		}
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)

		w := i
		go func() {
			defer wg.Done()
			writer(w)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			reader()
		}()
	}
	wg.Wait()
}
