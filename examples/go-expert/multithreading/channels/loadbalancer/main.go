package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for d := range data {
		fmt.Printf("worker %d received %d\n", workerId, d)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workers := 10

	for i := range workers {
		go worker(i, data)
	}

	for i := range 100 {
		data <- i
	}
}
