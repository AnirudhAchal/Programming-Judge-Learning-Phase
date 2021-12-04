package main

import (
	"fmt"
	"time"
)

var i int = 0

func diner(arr [5]int) {
	for i := 0; i < 5; i++ {
		if arr[i] == 1 && arr[(i+1)%5] == 1 {
			arr[i] = 0
			arr[(i+1)%5] = 0
			fmt.Printf("eating %d\n", i)
			arr[i] = 1
			arr[(i+1)%5] = 1
			fmt.Printf("thinking %d\n", i)
		}

	}
}

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = 1
	}
	go diner(arr)
	go diner(arr)
	go diner(arr)
	go diner(arr)
	time.Sleep(10 * time.Millisecond)
}
