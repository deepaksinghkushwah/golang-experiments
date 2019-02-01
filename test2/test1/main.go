package main

import "fmt"

func main() {
	a := [...]int{10, 20}
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	b := []int{10, 20}
	fmt.Printf("%T\n", b)
	fmt.Println(b)
}
