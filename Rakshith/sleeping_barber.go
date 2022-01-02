package main

import (
	"fmt"
	"sync"
	"time"
)

var CHECKING_TIME = 100
var WAITING_ROOM_SIZE = 5
var wg *sync.WaitGroup
var NUM_CUSTOMERS = 20

const (
	sleeping = iota
	cutting
	checking
)

type Barber struct{
	state int
	mu sync.Mutex
}

type WaitingRoom struct {
	arr []int
	mu sync.Mutex
}

/**

Barber goroutine 

Lock barber and waiting room
check if waiting room has customers
	yes => unlock room and barber, pick the first customer and hair cut
	no => unlock room and barber, go to sleep and wait for customer to wake up
*/

func barber(barber *Barber, customer chan int, room *WaitingRoom){
	for{
		barber.mu.Lock()
		room.mu.Lock()
		barber.state = checking
		fmt.Println("Barber checking waiting room")
		time.Sleep(time.Millisecond * 200)
		
		if(len(room.arr) > 0){
			c := room.arr[0]
			room.arr = room.arr[1:]
			room.mu.Unlock()
			hair_cut(barber, c)	
			barber.mu.Unlock()	
		} else{			
			room.mu.Unlock()
			barber.mu.Unlock()
			select {
			case c := <- customer :
				barber.state = cutting
				fmt.Printf("Customer %d waking up barber\n", c)				
				hair_cut(barber, c)
			default :
				barber.state = sleeping
				fmt.Println("Barber : No one's here I'll sleep now")
				c := <- customer
				barber.state = cutting
				hair_cut(barber, c)	
			}
		}		
	}
}
/**

Customer goroutine 

Lock barber 
check his status 
	sleeping => wake up barber, unlock barber
	cutting =>  unlock barber
				lock waiting room
				check waiting room size
					full => leave shop
					not full => sit at the end of queue
				unlock waiting room
*/
func customer(barber *Barber, customer chan int, c int, room *WaitingRoom){
	
	time.Sleep(time.Millisecond * 100)
	barber.mu.Lock()
	fmt.Printf("Customer %d checking barber room\n", c)
	if barber.state == cutting{
		barber.mu.Unlock()
		room.mu.Lock()
		defer room.mu.Unlock()
		if len(room.arr) == WAITING_ROOM_SIZE {
			fmt.Printf("Customer %d has no space in waiting room. Leaving shop..\n", c)
			wg.Done()
		}else{
			fmt.Printf("Customer %d sits in the waiting room\n", c)
			room.arr = append(room.arr, c)
		}
	} else if barber.state == sleeping {
		fmt.Printf("Customer %d waking up barber\n", c)
		customer <- c
		barber.mu.Unlock()
	}
}

func hair_cut(barber *Barber, customer int){
	fmt.Printf("Customer %d getting a hair cut\n", customer)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Customer %d had a nice haircut\n", customer)
	wg.Done()
}

func main(){

	wg = new(sync.WaitGroup)
	b := new(Barber)
	customer_chan := make(chan int)
	room := new(WaitingRoom)
	go barber(b, customer_chan, room)
	time.Sleep(time.Millisecond * 100) // let the barber get ready
	for i := 1; i <= NUM_CUSTOMERS; i++ {	
		wg.Add(1)
		go customer(b, customer_chan, i, room)
		// time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("All customers done")
}