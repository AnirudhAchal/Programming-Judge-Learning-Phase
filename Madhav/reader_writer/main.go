package main

import (
	"fmt"
	"sync"
	"time"
)

type a struct {
	sync.RWMutex
	i int
}

var j int = 5

func writer(v *a) {
	v.Lock()
	fmt.Println("writing started")
	v.i = j
	j++
	fmt.Println("writing finished")
	v.Unlock()
}

func reader(v *a) {
	for i := 0; i < 10; i++ {
		v.RLock()
		fmt.Println("reading started")
		fmt.Println(v.i)
		fmt.Println("reading finished")
		v.RUnlock()
	}
}

func main() {
	var v a
	go writer(&v)
	go reader(&v)
	go writer(&v)
	go reader(&v)
	go writer(&v)
	go reader(&v)
	time.Sleep(10 * time.Millisecond)
}
