package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
)

func main() {
	var i = 1
	file, err := os.OpenFile("./file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	for i < 10 {
		cstr, _ := uuid.NewV4()
		str := fmt.Sprintf("Hello World: %s \n", cstr.String())
		file.WriteString(str)
		i++
	}

	c, e := ioutil.ReadFile("file.txt")
	if e != nil {
		log.Fatalln(e)
	}

	fmt.Println(string(c))

}
