package main

import (
	"fmt"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
)

// Person struct for export globally
type Person struct {
	Name string
	Age  int
}

func fullname(p Person) {
	fmt.Println(p.Name)
}

func main() {
	var list []Person
	const x = 20
	//list = append(list, Person{Name: "Deepak", Age: 20})
	for i := 1; i <= 10; i++ {
		pl := Person{Name: fmt.Sprintf("Test %d", i), Age: 20 * i}
		list = append(list, pl)
		fullname(pl)
	}

	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	checkErr(err)
	defer file.Close()

	for _, item := range list {
		uuid, _ := uuid.NewV4()
		file.WriteString(item.Name + " " + uuid.String() + "\n")
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x)
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
