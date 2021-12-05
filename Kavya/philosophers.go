package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var forks [5]sync.Mutex
var philDone [5]bool
var avail [5]int

func checkAvail(i int) bool {
	var right int = i
	var left int = (i + 1) % 5

	//check if both forks are available
	if avail[left] == 1 && avail[right] == 1 {
		return true
	}
	return false
}
func dine(i int) {
	mu.Lock()

	if checkAvail(i) {
		avail[i] = 0
		avail[(i+1)%5] = 0

		forks[i].Lock()
		forks[(i+1)%5].Lock()

		fmt.Printf("Philosopher %v dining\n", i+1)
		time.Sleep(100 * time.Millisecond) //dining

		forks[i].Unlock()
		forks[(i+1)%5].Unlock()

		avail[i] = 1
		avail[(i+1)%5] = 1
		philDone[i] = true
		fmt.Printf("Philosopher %v done\n", i+1)

		mu.Unlock()
	} else {
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		avail[i] = 1
		philDone[i] = false
	}

	for {
		for i := 0; i < 5; i++ {
			if philDone[i] == false {
				i := i
				wg.Add(1)
				go func() {
					defer wg.Done()
					dine(i)
				}()
			}
		}
		wg.Wait()

		check := 0
		for i := 0; i < 5; i++ {
			if philDone[i] == true {
				check += 1
			}
		}

		if check == 5 {
			break
		}
	}
}
