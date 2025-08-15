package main

import "fmt"

// func Filter(items []int, predicate func(int) bool) []int {
// 	var result []int
// 	for _, item := range items {
// 		if predicate(item) {
// 			result = append(result, item)
// 		}
// 	}
// 	return result
// }

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("even numbers:", evenNumbers)

	names := []string{"Alice", "Bob", "Charlie", "David", "Anna"}

	liNames := Filter(names, func(s string) bool {
		return len(s) >= 2 && s[len(s)-2:] == "li"
	})
	fmt.Println("names ending with 'li':", liNames)

	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	youngPeople := Filter(people, func(p Person) bool {
		return p.Age < 30
	})
	fmt.Println("young people:", youngPeople)
}
