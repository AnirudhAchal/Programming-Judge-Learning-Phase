package main

import (
	"fmt"
	"sync"
	"time"
)

type chopstick struct {
	mu sync.Mutex
}

type philosopher struct {
	num         int
	left, right *chopstick
}

func eat(p philosopher, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Preparing to eat : %d\n", p.num)
	p.left.mu.Lock()
	p.right.mu.Lock()
	fmt.Printf("Philosopher %d got the chopsticks. Started eating!\n", p.num)
	time.Sleep(time.Second)
	fmt.Printf("Philosopher %d finished eating\n", p.num)
	p.left.mu.Unlock()
	p.right.mu.Unlock()
}

func main() {

	count := 5

	wg := new(sync.WaitGroup)

	chopsticks := make([]*chopstick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = new(chopstick)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		philosophers[i] = &philosopher{i, chopsticks[i], chopsticks[(i+1)%count]}

		go eat(*philosophers[i], wg)
	}
	wg.Wait()
}
