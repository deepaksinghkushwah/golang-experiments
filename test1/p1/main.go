package main

import (
	"fmt"

	"github.com/deepaksinghkushwah/test1/stringutils"
)

func main() {
	fmt.Println(stringutils.MyName)
	//fmt.Println(stringutils.Reverse(stringutils.MyName))
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println(name)
}
