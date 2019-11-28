package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	n := average(data...)
	fmt.Println(n)
	greetmaker := greet()
	fmt.Println(greetmaker())
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func greet() func() string {
	return func() string {
		return "Helo World"
	}
}
