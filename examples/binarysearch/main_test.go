package main

import (
	"fmt"
	"testing"
)

type TestCaseOne struct {
	items    []int
	target   int
	expected int
}

var testCasesOne = []TestCaseOne{
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
	for _, tc := range testCasesOne {
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
	for _, tc := range testCasesOne {
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

func largeTest() *TestCaseOne {
	size := 10000000
	numbers := make([]int, 0, size)
	for i := size; i > 0; i-- {
		numbers = append(numbers, i)
	}
	return &TestCaseOne{
		items:    numbers,
		target:   2,
		expected: 9999998,
	}
}

// go test -bench=.
func BenchmarkLocateCardLinear(b *testing.B) {
	test := largeTest()
	for b.Loop() {
		LocateCardLinear(test.items, test.target)
	}
}

func BenchmarkLocateCard(b *testing.B) {
	test := largeTest()
	for b.Loop() {
		LocateCard(test.items, test.target)
	}
}

type TestCaseTwo struct {
	items    []int
	target   int
	expected Tuple
}

var testCasesTwo = []TestCaseTwo{
	{
		items:    []int{},
		target:   1,
		expected: Tuple{Left: -1, Right: -1},
	},
	{
		items:    []int{1, 2, 4, 4, 5},
		target:   6,
		expected: Tuple{Left: -1, Right: -1},
	},
	{
		items:    []int{1},
		target:   1,
		expected: Tuple{Left: 0, Right: 0},
	},
	{
		items:    []int{1, 2, 3, 4, 5},
		target:   4,
		expected: Tuple{Left: 3, Right: 3},
	},
	{
		items:    []int{1, 2, 4, 4, 4, 4, 4, 4, 5},
		target:   4,
		expected: Tuple{Left: 2, Right: 7},
	},
	{
		items:    []int{-1, -2, -2, -2, 4, 5, 6, 7},
		target:   -2,
		expected: Tuple{Left: 1, Right: 3},
	},
}

func TestFirstAndLastPosition(t *testing.T) {
	for _, tc := range testCasesTwo {
		testName := fmt.Sprintf("%d, %d", tc.items, tc.target)
		t.Run(testName, func(t *testing.T) {
			// act
			result := FirstAndLastPosition(tc.items, tc.target)

			// assert
			if result != tc.expected {
				t.Errorf("got %d, expected %d", result, tc.expected)
			}
		})
	}
}
