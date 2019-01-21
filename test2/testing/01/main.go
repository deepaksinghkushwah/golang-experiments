package main

import "fmt"

func main() {
	x := []int{1, 2, 3}

	fmt.Println(mySum(x...))
}

func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
