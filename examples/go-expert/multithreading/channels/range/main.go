package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan int)

	// Thread 2
	go publish(channel)

	reader(channel)
	// for r := range channel {
	// 	fmt.Printf("received: %d\n", r)
	// }
}

func reader(channel chan int) {
	for r := range channel {
		fmt.Printf("received: %d\n", r)
	}
}

func publish(channel chan int) {
	for i := range 5 {
		channel <- i
	}
	close(channel)
}
