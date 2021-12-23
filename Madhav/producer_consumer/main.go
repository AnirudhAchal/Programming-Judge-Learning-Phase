package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

func producer_action(c *chan int, wg1 *sync.WaitGroup) {
	mu.Lock()
	fmt.Println("Producing started")
	if len(*c) == 5 {
		fmt.Println("Producing finished")
		mu.Unlock()
	} else {
		j := (rand.Intn(100) % 5) + 1
		fmt.Println(j)
		*c <- j
		fmt.Println("Producing finished")
		mu.Unlock()
	}
	wg1.Done()
}

func consumer_action(c *chan int, wg1 *sync.WaitGroup) {
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
	wg1.Done()
}

func producers(j int, c *chan int, wg *sync.WaitGroup) {
	var wg1 sync.WaitGroup
	for i := 0; i < j; i++ {
		x := (rand.Intn(100) % 5) + 1
		sec := time.Duration(x) * time.Millisecond
		time.Sleep(sec)
		wg1.Add(1)
		go producer_action(c, &wg1)
	}
	wg1.Wait()
	wg.Done()
}

func consumers(j int, c *chan int, wg *sync.WaitGroup) {
	var wg1 sync.WaitGroup
	for i := 0; i < j; i++ {
		x := (rand.Intn(100) % 5) + 1
		sec := time.Duration(x) * time.Millisecond
		time.Sleep(sec)
		wg1.Add(1)
		go consumer_action(c, &wg1)
	}
	wg1.Wait()
	wg.Done()
}

func main() {
	c := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go producers(10, &c, &wg)
	wg.Add(1)
	go consumers(10, &c, &wg)
	wg.Wait()
}
