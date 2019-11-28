package main

import "fmt"

type person struct {
	first string
	last  string
}

type secrateAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("Speaking as person", p.first, p.last)
}

func (sa secrateAgent) speak() {
	fmt.Println("Speaking as secrateAgent", sa.person.first, sa.person.last)
}

type human interface {
	speak()
}

func interfaceFunc(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am interface", h.(person).first)
	case secrateAgent:
		fmt.Println("I am interface", h.(secrateAgent).first)
	}
	//fmt.Println("I am human", h)
}

func main() {
	sa1 := secrateAgent{
		person: person{
			first: "Deepak",
			last:  "Kushwah",
		},
		ltk: true,
	}

	sa2 := secrateAgent{
		person: person{
			first: "Kanchan",
			last:  "Kushwah",
		},
		ltk: false,
	}

	p1 := person{
		first: "Yaksh",
		last:  "Kushwah",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

	interfaceFunc(sa1)
	interfaceFunc(sa2)
	interfaceFunc(p1)

}
