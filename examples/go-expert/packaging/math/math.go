package math

// type Math struct {
// 	A int
// 	B int
// }

// func (m Math) Add() int {
// 	return m.A + m.B
// }

type Math struct {
	a int
	b int
}

func NewMath(a, b int) Math {
	return Math{a: a, b: b}
}

func (m Math) Add() int {
	return m.a + m.b
}
