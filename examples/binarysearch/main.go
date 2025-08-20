/*
Alice has some cards with numbers written on them.
She arranges the cards in decreasing order, and lays them out face down in a sequence on a table.
She challenges Bob to pick out the card containing a given number by turning over as few cards as possible.
Write a function to help Bob locate the card.
*/

package main

import "fmt"

func LocateCardLinear(cards []int, query int) int {
	position := 0
	for position < len(cards) {
		if cards[position] == query {
			return position
		}
		position += 1
	}
	return -1
}

// complex
// func LocateCard(cards []int, query int) int {
// 	left, right := 0, len(cards)-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if cards[mid] == query {
// 			if mid > 0 && cards[mid-1] == query {
// 				right = mid - 1
// 			} else {
// 				return mid
// 			}
// 		} else {
// 			if cards[mid] < query {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 	}
// 	return -1
// }

func binarySearch(left, right int, condition func(int) string) int {
	for left <= right {
		mid := left + (right-left)/2
		result := condition(mid)
		if result == "found" {
			return mid
		}
		if result == "left" {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func LocateCard(cards []int, query int) int {
	left, right := 0, len(cards)-1
	condition := func(mid int) string {
		if cards[mid] == query {
			if mid > 0 && cards[mid-1] == query {
				return "left"
			} else {
				return "found"
			}
		}
		if cards[mid] > query {
			return "right"
		} else {
			return "left"
		}
	}
	return binarySearch(left, right, condition)
}

/*
Question: Given an array of integers nums sorted in ascending order, find the starting and ending position of a given number.
This differs from the problem in only two significant ways:
  1 - The numbers are sorted in increasing order.
  2 -  We are looking for both the increasing order and the decreasing order.
Here's the full code for solving the question, obtained by making minor modifications to our previous function:
*/

func FirstPosition(items []int, target int) int {
	left, right := 0, len(items)-1
	condition := func(mid int) string {
		if items[mid] == target {
			if mid > 0 && items[mid-1] == target {
				return "left"
			} else {
				return "found"
			}
		}
		if items[mid] < target {
			return "right"
		} else {
			return "left"
		}
	}
	return binarySearch(left, right, condition)
}

func LastPosition(items []int, target int) int {
	left, right := 0, len(items)-1
	condition := func(mid int) string {
		if items[mid] == target {
			if mid < len(items)-1 && items[mid+1] == target {
				return "right"
			} else {
				return "found"
			}
		}
		if items[mid] < target {
			return "right"
		} else {
			return "left"
		}
	}
	return binarySearch(left, right, condition)
}

type Tuple struct {
	Left  int
	Right int
}

func FirstAndLastPosition(items []int, target int) Tuple {
	return Tuple{
		Left:  FirstPosition(items, target),
		Right: LastPosition(items, target),
	}
}

/*
You are given list of numbers, obtained by rotating a sorted list an unknown number of times. Write a function to determine the minimum number of times the original sorted list was rotated to obtain the given list. Your function should have the worst-case complexity of O(log N), where N is the length of the list. You can assume that all the numbers in the list are unique.

Example: The list [5, 6, 9, 0, 2, 3, 4] was obtained by rotating the sorted list [0, 2, 3, 4, 5, 6, 9] 3 times.

We define "rotating a list" as removing the last element of the list and adding it before the first element. E.g. rotating the list [3, 2, 4, 1] produces [1, 3, 2, 4].

"Sorted list" refers to a list where the elements are arranged in the increasing order e.g. [1, 3, 5, 7].
*/

func CountRotationsLinear(items []int) int {
	// manual index, prone to off-by-one errors
	// position := 0
	// for position <= len(items)-1 {
	// 	if position > 0 && items[position] < items[position-1] {
	// 		return position
	// 	}
	// 	position++ // manual increment
	// }

	// Go-idiomatic, safe, and clean
	for position := range items {
		if position > 0 && items[position] < items[position-1] {
			return position
		}
	}
	return -1
}

func main() {
	// cards := []int{13, 11, 10, 7, 4, 3, 1, 0}
	// query := 1
	// fmt.Println(LocateCardLinear(cards, query))

	// items := []int{1, 2, 4, 4, 4, 4, 5}
	// target := 4
	// fmt.Println(FirstAndLastPosition(items, target))

	items := []int{5, 6, 9, 0, 2, 3, 4}
	fmt.Println(CountRotationsLinear(items))
}
