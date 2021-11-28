package main

import (
	"fmt"
	"time"
)

type queue struct {
	tail int
	head int
	data [10]int
}

func in(q *queue) {
	q.head = 0
	q.tail = 0
}

func producer(q *queue) {
	i := 0
	for true {
		for ((q.tail + 1) % 10) == q.head {
		}
		q.data[q.tail] = i
		fmt.Printf("Inserted %d \n", i)
		i++
		q.tail = (q.tail + 1) % 10
	}
}

func consumer(q *queue) {
	a := 0
	for true {
		for q.head == q.tail {
		}
		a = q.data[q.head]
		fmt.Printf("Extracted %d \n", a)
		q.head = (q.head + 1) % 10
	}

}

func main() {
	var q queue
	q.head = 0
	q.tail = 0
	go producer(&q)
	go consumer(&q)
	time.Sleep(10 * time.Millisecond)
}
