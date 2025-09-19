package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(5)

	// Thread 2
	go publish(channel)

	// Thread 3
	go reader(channel, &wg)

	wg.Wait()
}

func reader(channel chan int, wg *sync.WaitGroup) {
	for r := range channel {
		fmt.Printf("received %d\n", r)
		wg.Done()
	}
}

func publish(channel chan int) {
	for i := range 5 {
		channel <- i
	}
	close(channel)
}
