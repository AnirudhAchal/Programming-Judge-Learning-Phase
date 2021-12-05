package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Unit struct {
	x int
	l sync.RWMutex
}

type ReadCount struct {
	x int
	l sync.Mutex
}

var u Unit
var r ReadCount

func write() {
	u.l.Lock()
	v := rand.Int() / 1000
	fmt.Println("Replacing", u.x, " with", v)
	u.x = v
	u.l.Unlock()
}

func read(p int) {
	u.l.RLock()
	r.l.Lock()
	r.x += 1
	r.l.Unlock()
	fmt.Println("Reader", p, "is reading:", u.x)
	fmt.Println("How many are reading ?", r.x)
	r.l.Lock()
	r.x -= 1
	r.l.Unlock()
	u.l.RUnlock()
}

func main() {
	n := 20
	u.x = 0
	for j := 0; j < n; j++ {
		go write()
		time.Sleep(time.Millisecond)
		go read(j)
		go read(2 * j)
		go read(3 * j)
	}
	time.Sleep(5 * time.Second)
}
