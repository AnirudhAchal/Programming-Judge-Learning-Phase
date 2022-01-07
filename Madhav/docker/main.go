package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func diner(arr [5]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		mu.Lock()
		if arr[i] == 1 && arr[(i+1)%5] == 1 {
			x := (rand.Intn(100) % 5) + 1
			sec := time.Duration(x) * time.Millisecond
			arr[i] = 0
			arr[(i+1)%5] = 0
			mu.Unlock()
			fmt.Printf("eating %d\n", i)
			time.Sleep(sec)
			mu.Lock()
			arr[i] = 1
			arr[(i+1)%5] = 1
			fmt.Printf("thinking %d\n", i)
			mu.Unlock()
		} else {
			mu.Unlock()
		}
	}
	wg.Done()
}

func main() {
	var arr [5]int
	var mu sync.Mutex
	for i := 0; i < 5; i++ {
		arr[i] = 1
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go diner(arr, &mu, &wg)
	wg.Add(1)
	go diner(arr, &mu, &wg)
	wg.Add(1)
	go diner(arr, &mu, &wg)
	wg.Add(1)
	go diner(arr, &mu, &wg)
	wg.Add(1)
	go diner(arr, &mu, &wg)
	wg.Wait()
}
