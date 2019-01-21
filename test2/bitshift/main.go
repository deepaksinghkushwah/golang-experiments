package main

import "fmt"

func main() {
	var i uint
	for i = 10; i < 40; i += 10 {
		fmt.Printf("i \t\t %d \t\t %b\n", i, 1<<i)
	}
}
