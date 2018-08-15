package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./content.txt")
	if err != nil {
		log.Fatalln(err)
	}
	str := ""
	split := strings.Split(string(dat), "\n")
	for _, item := range split {
		if len(item) > 0 {
			nextsplit := strings.Split(item, "Pindi: ")
			if len(nextsplit) > 1 {
				str += nextsplit[1] + "\n"
			}

		}

	}
	fmt.Println(str)
}
