package main

import (
	"fmt"
	"sync"
	"time"
)

func printSorted(val int) {
	time.Sleep(time.Duration(val) * time.Millisecond)
	fmt.Printf("%d ", val)
}
func main() {
	arr := [5]int{3, 6, 4, 2, 5}
	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			printSorted(v)
		}(v)
	}
	wg.Wait()
}
