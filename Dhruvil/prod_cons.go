package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mu sync.Mutex

var count = 0
var in = 0
var out = 0
var data [10]int

func producer(wg *sync.WaitGroup) {

	defer wg.Done()
	for true {
		mu.Lock()
		if count == 10 {
			mu.Unlock()
		} else {
			count++
			data[in] = rand.Int()
			fmt.Printf("Produced %d\n", data[in])
			in = (in + 1) % 10
			mu.Unlock()
		}
	}
}
func consumer(wg *sync.WaitGroup) {

	defer wg.Done()
	for true {
		mu.Lock()
		if count == 0 {
			mu.Unlock()
		} else {
			count--
			fmt.Printf("Consumed %d\n", data[out])
			out = (out + 1) % 10
			mu.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(&wg)
	go consumer(&wg)
	wg.Wait()

	// time.Sleep(2 * time.Millisecond)

}
