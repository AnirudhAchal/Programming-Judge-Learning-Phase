package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var chair sync.Mutex
var wg sync.WaitGroup

var barberStatus int //Sleeping, cutting
var waitingRoom chan int = make(chan int, 3)

func barber() {
	for {
		mu.Lock()

		select {
		case c := <-waitingRoom:
			{
				barberStatus = 1
				chair.Lock()
				mu.Unlock()

				fmt.Printf("Customer %v gets haircut\n", c)
				time.Sleep(time.Microsecond)
				chair.Unlock()
				fmt.Printf("Customer %v is happy\n", c)
				wg.Done()

			}
		default:
			mu.Unlock()
			barberStatus = 0
			fmt.Printf("Time to sleep!\n")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func customer(i int) {

	mu.Lock()
	if barberStatus == 0 {
		barberStatus = 1
		fmt.Printf("Customer %v wakes up barber\n", i)
		chair.Lock()
		mu.Unlock()

		fmt.Printf("Customer %v gets haircut\n", i)
		time.Sleep(time.Microsecond)
		chair.Unlock()
		fmt.Printf("Customer %v is happy\n", i)
		wg.Done()

		time.Sleep(100 * time.Millisecond)
	} else {
		if len(waitingRoom) == cap(waitingRoom) {
			fmt.Printf("Customer %v leaving.\n", i)
			wg.Done()
		} else {
			waitingRoom <- i
			fmt.Printf("Customer %v sat in waiting room\n", i)
		}
		mu.Unlock()
	}
}
func main() {
	barberStatus = 0

	go barber()

	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go customer(i)
	}
	wg.Wait()

	fmt.Printf("All customers are done, time to close.\n")
}
