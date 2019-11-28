package main

import "fmt"

type person struct {
	id   int
	name string
}

func main() {
	list := map[string]int{
		"Deepak":  29,
		"Kanchan": 32,
	}

	if v, ok := list["yaksh"]; !ok {
		fmt.Println("if cond ", v)
	}

	s := make(map[int]person)

	s[0] = person{id: 1, name: "Deepak"}
	s[1] = person{id: 2, name: "Kanchan"}
	fmt.Println(s)

	for _, v := range s {
		fmt.Println(v.id, v.name)
	}

	delete(s, 0)
	fmt.Println(s)
}
