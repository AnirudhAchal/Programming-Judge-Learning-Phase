package main

import (
	"fmt"
)

func producer(c chan int, finished chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Sending %d to channel\n", i)
		c <- i
		fmt.Printf("Sent %d to channel\n", i)
	}
	finished <- true
}

func consumer(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("\t\t\t\tWaiting for %d\n", i)
		fmt.Println("\t\t\t\tConsumed ", <-c)
	}

}

func main() {

	c := make(chan int)
	finished := make(chan bool)
	go producer(c, finished)
	go consumer(c)
	select {
	case <-finished:
		return
	}

	/**
	Remarks
	Consumed 'i' comes after 'sending 'i' to channel'
	But consumed 'i' may come before 'sent 'i' to channel'
	*/

}
