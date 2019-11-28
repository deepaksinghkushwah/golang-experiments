package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type animal struct {
	ID   int
	Name string
	Eat  string
	Mood []string
}

func main() {
	s := animal{
		ID:   1,
		Name: "Dog",
		Eat:  "bone",
		Mood: []string{"game", "music", "roaming"},
	}
	ss, err := json.Marshal(s)
	checkErr(err)
	os.Stdout.Write(ss)
	fmt.Println()
	fmt.Println("-----------------------")
	var x animal
	err = json.Unmarshal(ss, &x)
	checkErr(err)
	fmt.Printf("%+v", x)
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
