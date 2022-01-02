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

func producer(index int, myArr *MyArray, wg *sync.WaitGroup) {
	defer wg.Done()
	myArr.mu.Lock()
	if len(myArr.arr) == MAX_SIZE {
		fmt.Printf("Array if full ! Producer %d cannot add\n", index)
		myArr.mu.Unlock()
	} else {
		fmt.Printf("Producing %d\n", index)
		myArr.arr = append(myArr.arr, index)
		myArr.mu.Unlock()
	}
}

func consumer(index int, myArr *MyArray, wg *sync.WaitGroup) {
	defer wg.Done()
	myArr.mu.Lock()
	if len(myArr.arr) == 0 {
		fmt.Printf("Consumer %d has nothing to consume\n", index)
		myArr.mu.Unlock()
	} else {
		x := myArr.arr[len(myArr.arr)-1]
		fmt.Printf("Consumed %d from consumer %d\n", x, index)
		myArr.arr = myArr.arr[:len(myArr.arr)-1]
		myArr.mu.Unlock()
	}

}

func main() {

	wg := new(sync.WaitGroup)
	myArray := MyArray{}
	for i := 1; i <= NUM_PRODUCERS; i++ {
		wg.Add(1)
		go producer(i, &myArray, wg)
	}
	for i := 1; i <= NUM_CONSUMERS; i++ {
		wg.Add(1)
		go consumer(i, &myArray, wg)

	}
	wg.Wait()

}
