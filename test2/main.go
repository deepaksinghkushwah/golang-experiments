package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	sort.Ints(x)
	lowest := x[0]
	highest := x[len(x)-1]
	fmt.Println("Lowest:", lowest, "Highest:", highest)
	fmt.Println("Capacity of X:", cap(x))
}
