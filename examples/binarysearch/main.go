/*
Alice has some cards with numbers written on them.
She arranges the cards in decreasing order, and lays them out face down in a sequence on a table.
She challenges Bob to pick out the card containing a given number by turning over as few cards as possible.
Write a function to help Bob locate the card.
*/

package main

import (
	"fmt"
)

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

func LocateCard(cards []int, query int) int {
	left, right := 0, len(cards)-1
	for left <= right {
		mid := left + (right-left)/2
		if cards[mid] == query {
			return mid
		}
		if cards[mid] < query {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	cards := []int{13, 11, 10, 7, 4, 3, 1, 0}
	query := 1
	fmt.Println(LocateCardLinear(cards, query))
}
