package main

import "fmt"

func main() {
	// init map
	// m := make(map[string]int)

	// m := map[string]int{}

	incomes := map[string]int{
		"John": 1000,
		"Jane": 2000,
		"Boby": 3000,
	}
	// fmt.Println(incomes)
	// fmt.Println(incomes["Boby"])

	// delete(incomes, "Boby")

	// fmt.Println(incomes)

	// incomes["Jack"] = 4000

	// fmt.Println(incomes)

	for name, income := range incomes {
		fmt.Printf("%s, %d\n", name, income)
	}

	for _, income := range incomes {
		fmt.Printf("%d\n", income)
	}
}
