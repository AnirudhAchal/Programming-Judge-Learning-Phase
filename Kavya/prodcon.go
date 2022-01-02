package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ch chan int = make(chan int, 4)
var mu sync.Mutex

func producer(w *sync.WaitGroup, p int) {
	w.Add(1)
	defer w.Done()
	mu.Lock()

	if len(ch) == cap(ch) {
		mu.Unlock()
	} else {
		x := rand.Intn(100)
		fmt.Printf("Producer %v sent %v\n", p, x)
		ch <- x
		mu.Unlock()
	}
	time.Sleep(5 * time.Millisecond)
}

func consumer(w *sync.WaitGroup, c int) {
	w.Add(1)
	defer w.Done()

	mu.Lock()

	if len(ch) == 0 {
		mu.Unlock()
	} else {
		v := <-ch
		fmt.Printf("Consumer %v received %v \n", c, v)
		mu.Unlock()
	}
	time.Sleep(time.Millisecond)
}
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		go producer(&wg, i)
	}

	for i := 1; i <= 5; i++ {
		go consumer(&wg, i)
	}

	wg.Wait()
}
