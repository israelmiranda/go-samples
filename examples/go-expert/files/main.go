package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write
	// size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("file created! size: %d bytes\n", size)

	// read
	fb, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fb))

	// read file buffer
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remove
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
