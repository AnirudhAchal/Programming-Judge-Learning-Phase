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
	v.i = j
	j++
	v.Unlock()
}

func reader(v *a) {
	for i := 0; i < 10; i++ {
		v.RLock()
		fmt.Println(v.i)
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
