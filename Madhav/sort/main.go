package main

import (
	"fmt"
	"time"
)

func print(a int) {
	time.Sleep(time.Duration(a) * time.Millisecond)
	fmt.Println(a)
}

func main() {
	arr := []int{5, 9, 2, 7, 3}
	for i := 0; i < len(arr); i++ {
		go print(arr[i])
	}
	time.Sleep(15 * time.Millisecond)
}
