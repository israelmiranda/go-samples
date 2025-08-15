package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	items    []int
	target   int
	expected int
}

var testCases = []TestCase{
	{
		items:    []int{13, 11, 10, 7, 4, 3, 1, 0},
		target:   1,
		expected: 6,
	},
	{
		items:    []int{4, 2, 1, -1},
		target:   4,
		expected: 0,
	},
	{
		items:    []int{3, -1, -9, -127},
		target:   -127,
		expected: 3,
	},
	{
		items:    []int{6},
		target:   6,
		expected: 0,
	},
	{
		items:    []int{9, 7, 5, 2, -9},
		target:   4,
		expected: -1,
	},
	{
		items:    []int{},
		target:   7,
		expected: -1,
	},
	{
		items:    []int{8, 8, 6, 6, 6, 6, 6, 6, 3, 2, 2, 2, 0, 0, 0},
		target:   6,
		expected: 2,
	},
}

// go test -v
func TestLocateCardLinearBasic(t *testing.T) {
	// arrange
	cards := []int{13, 11, 10, 7, 4, 3, 1, 0}
	query := 1

	// act
	result := LocateCardLinear(cards, query)

	// assert
	if result != 6 {
		t.Errorf("LocateCardLinear(cards, query) = %d; want 6", result)
	}
}

func TestLocateCardLinearTableDriven(t *testing.T) {
	for _, tc := range testCases {
		testName := fmt.Sprintf("%d, %d", tc.items, tc.target)
		t.Run(testName, func(t *testing.T) {
			// act
			result := LocateCardLinear(tc.items, tc.target)

			// assert
			if result != tc.expected {
				t.Errorf("got %d, expected %d", result, tc.expected)
			}
		})
	}
}

func TestLocateCard(t *testing.T) {
	for _, tc := range testCases {
		testName := fmt.Sprintf("%d, %d", tc.items, tc.target)
		t.Run(testName, func(t *testing.T) {
			// act
			result := LocateCard(tc.items, tc.target)

			// assert
			if result != tc.expected {
				t.Errorf("got %d, expected %d", result, tc.expected)
			}
		})
	}
}

// go test -bench=.
func BenchmarkLocateCardLinear(b *testing.B) {
	cards := []int{13, 11, 10, 7, 4, 3, 1, 0}
	query := 0
	for b.Loop() {
		LocateCardLinear(cards, query)
	}
}

func BenchmarkLocateCard(b *testing.B) {
	cards := []int{13, 11, 10, 7, 4, 3, 1, 0}
	query := 0
	for b.Loop() {
		LocateCard(cards, query)
	}
}
