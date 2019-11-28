package main

import "fmt"

func main() {
	a := 10
	b := &a
	testme(b)
}

func testme(x *int) {
	fmt.Println(*x)
}
