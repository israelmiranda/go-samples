package main

// Thread 1
func main() {
	forever := make(chan bool)

	// Thread 2
	go func() {
		for i := range 5 {
			println(i)
		}
		forever <- true
	}()

	<-forever
}
