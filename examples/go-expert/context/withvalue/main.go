package main

import (
	"context"
	"fmt"
)

// Define a custom, unexported type for the key.
type contextKey string

// Use a unique key for the context.
const myKey contextKey = "myKey"

func main() {
	// SA1029 warning here
	ctx := context.WithValue(context.Background(), myKey, "myValue")

	value := ctx.Value(myKey)
	fmt.Println(value)
}
