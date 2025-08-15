package main

import "fmt"

// func Map(items []int, transform func(int) string) []string {
// 	result := make([]string, len(items))
// 	for i, item := range items {
// 		result[i] = transform(item)
// 	}
// 	return result
// }

func Map[T any, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = transform(item)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	stringNumbers := Map(numbers, func(n int) string {
		return fmt.Sprintf("Value: %d", n)
	})
	fmt.Println("ints to strings:", stringNumbers)

	names := []string{"Alice", "Bob", "Charlie"}

	nameLengths := Map(names, func(s string) int {
		return len(s)
	})
	fmt.Println("string lengths:", nameLengths)

	isEven := Map(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("is even:", isEven)

	type User struct {
		ID   int
		Name string
	}
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	userNames := Map(users, func(u User) string {
		return u.Name
	})
	fmt.Println("user names:", userNames)

	floatNumbers := Map(numbers, func(n int) float64 {
		return float64(n) * 1.5
	})
	fmt.Println("ints to float64:", floatNumbers)
}
