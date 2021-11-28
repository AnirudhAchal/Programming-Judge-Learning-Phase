package main

import (
	"fmt"
	"sync"
	"time"
)

func printSorted(val int) {
	time.Sleep(time.Duration(val) * time.Second)
	fmt.Print(val)
}

func main() {
	var arr []int = []int{4, 2, 3, 1, 5}

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)

		x := arr[i]
		go func() {
			defer wg.Done()
			printSorted(x)
		}()
		//go printSorted(arr[i])
	}
	wg.Wait()
}
