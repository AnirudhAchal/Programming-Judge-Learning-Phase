package main

import (
	"fmt"
	"time"
)

func producer(c chan int, n int) {
	for j := 0; j < n; j++ {
		fmt.Println("Producing the value:", j)
		c <- j
		time.Sleep(time.Millisecond)
	}
	
}

func consumer(c chan int, n int) {
	for j := 0; j < n; j++ {
		x := <-c
		fmt.Println("Consuming the value:", x)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	n := 1000
	c := make(chan int, n)
	go producer(c, n)
	go consumer(c, n)
	time.Sleep(10 * time.Second)
}
