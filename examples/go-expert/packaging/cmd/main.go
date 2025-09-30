package main

import (
	"fmt"

	"github.com/israelmiranda/go-samples/examples/go-expert/packaging/math"
)

func main() {
	// m := math.Math{A: 1, B: 2}
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
