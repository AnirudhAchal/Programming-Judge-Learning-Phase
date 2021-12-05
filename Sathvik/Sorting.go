package main

import (
	"fmt"
	"time"
)

func pushToChannel(channel chan int, a int) {
	time.Sleep(time.Duration(a) * time.Millisecond)
	channel <- a
}

func main() {
	var n int
	fmt.Println("We can sort a list of distinct positive integers for you.")
	fmt.Println("How many numbers do you have?")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Too few numbers!")
	} else {
		fmt.Println("Enter the list: ")
		channel := make(chan int, n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Scan(&a)
			go pushToChannel(channel, a)
		}
		fmt.Println("The sorted list is: ")
		for ; n > 0; n-- {
			fmt.Print(<-channel, " ")
		}
		fmt.Println()
	}

}
