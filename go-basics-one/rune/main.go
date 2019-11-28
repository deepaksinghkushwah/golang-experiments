package main

import "fmt"

func main() {
	/*for i := 49; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}*/
	switch "Archana" {
	case "Kanchan", "Deepak":
		fmt.Println("Found Deepak")
	case "Manish", "Archana":
		fmt.Println("Found manish and archana")
	default:
		fmt.Println("No one found")
	}
}
