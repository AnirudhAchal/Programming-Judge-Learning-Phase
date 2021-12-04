package main

import (
	"fmt"
	"sync"
)

var MAX_SIZE int = 5
var NUM_PRODUCERS int = 10
var NUM_CONSUMERS int = 10

type MyArray struct {
	mu  sync.Mutex
	arr []int
}

func producer(myArr *MyArray, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= NUM_PRODUCERS; i++ {
		myArr.mu.Lock()
		if len(myArr.arr) == MAX_SIZE {
			fmt.Printf("Array if full ! Producer %d cannot add\n", i)
			myArr.mu.Unlock()
		} else {
			fmt.Printf("Producing %d\n", i)
			myArr.arr = append(myArr.arr, i)
			myArr.mu.Unlock()
		}
	}
}

func consumer(myArr *MyArray, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= NUM_CONSUMERS; i++ {
		myArr.mu.Lock()
		if len(myArr.arr) == 0 {
			fmt.Printf("Consumer %d has nothing to consume\n", i)
			myArr.mu.Unlock()
		} else {
			x := myArr.arr[len(myArr.arr)-1]
			fmt.Printf("Consumed %d from consumer %d\n", x, i)
			myArr.arr = myArr.arr[:len(myArr.arr)-1]
			myArr.mu.Unlock()
		}
	}

}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(2)
	myArr := MyArray{}
	go producer(&myArr, wg)
	go consumer(&myArr, wg)
	wg.Wait()

}
