package main

import (
	"fmt"
	"sort"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

// ByAge sort by
type ByAge []Person

// ByName sort by
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	a := []int{2, 4, 1, 68, 5, 3}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

	b := []string{"Deepak", "Kanchan", "Dhairya", "Yaksh"}
	fmt.Println(b)
	sort.Strings(b)
	fmt.Println(b)

	p := []Person{
		{
			Name: "Deepak",
			Age:  54,
		},
		{
			Name: "Kanchan",
			Age:  40,
		},
	}

	fmt.Println(p)
	fmt.Println()
	fmt.Println("----------------------")
	fmt.Println()
	fmt.Println("Sort By Age")
	sort.Sort(ByAge(p))
	fmt.Println(p)

	fmt.Println("Sort By Name")
	sort.Sort(ByName(p))
	fmt.Println(p)
}
