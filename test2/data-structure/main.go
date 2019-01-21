package main

import (
	"fmt"
)

func main() {
	/*x := "Deepak Singh Kushwah"
	fmt.Println(x)
	for i, item := range x {
		fmt.Printf("%d \t\t %s\n", int32(i), string(item))
	}*/

	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(y)

	y = append(y[:1], y[4:]...)
	fmt.Println(y)

	jb := []string{"Deepak", "Kushwah", "strabery"}
	mp := []string{"Money", "Penny", "ginger"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
