package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var NUM_READERS int = 10
var NUM_WRITERS int = 10

type MyNum struct {
	mu sync.RWMutex
	x  int
}

func reader(myNum *MyNum, wg *sync.WaitGroup) {
	defer wg.Done()
	myNum.mu.RLock()
	defer myNum.mu.RUnlock()
	fmt.Printf("Reading value %d\n", myNum.x)
}

func writer(myNum *MyNum, wg *sync.WaitGroup) {
	defer wg.Done()
	myNum.mu.Lock()
	defer myNum.mu.Unlock()
	myNum.x = rand.Intn(100)
	fmt.Printf("Wrote value %d\n", myNum.x)
}

func main() {

	wg := new(sync.WaitGroup)
	myNum := MyNum{}
	for i := 1; i <= NUM_WRITERS; i++ {
		wg.Add(1)
		go writer(&myNum, wg)
	}
	for i := 1; i <= NUM_READERS; i++ {
		wg.Add(1)
		go reader(&myNum, wg)

	}
	wg.Wait()

}
