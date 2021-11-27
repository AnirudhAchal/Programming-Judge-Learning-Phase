package main

import (
	"fmt"
	"time"
)

func sort(s []int, c chan int) {
	for j := 0; j < len(s); j++ {
		go sleepAndPutToChannel(s[j], c)
	}
}

func sleepAndPutToChannel(n int, c chan int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
	c <- n
}

func main() {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	//Testing with random numbers
	// s := [100]int{}
	// const RANGE = 1000
	// for i := 0; i < len(s); i++ {
	// 	s[i] = rand.Intn(RANGE)
	// }

	c := make(chan int, len(s))
	sort(s, c)
	fmt.Println("Sorted nunmbers")
	for i := 0; i < len(s); i++ {
		fmt.Print(<-c, " ")
	}

}
