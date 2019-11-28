package main

import (
	"fmt"
)

type human struct {
	Name   string
	Age    int
	Gender string
}

type student struct {
	human
	school string
}

type employee struct {
	human
	company string
}

func (h *human) sayHi() {
	fmt.Println("Hello " + h.Name)
}

func (e *employee) sayHi() {
	fmt.Println("Hola " + e.Name)
}

func main() {
	deepak := student{human{Name: "Deepak Singh Kushwah", Age: 37, Gender: "Male"}, "MIT"}
	kanchan := employee{human{Name: "Kanchan", Age: 40, Gender: "Male"}, "BHEL"}
	fmt.Println(deepak)
	kanchan.sayHi()
}
