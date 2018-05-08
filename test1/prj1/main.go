/**
 * mail-ids.txt is simple txt file with new line seprated email ids
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// for array of slice
	ids := []string{"test1@gmail.com", "test2@gmail.com", "test3@gmail.com"}

	for index, id := range ids {
		newID := strings.Replace(id, "@gmail.com", "", 1)
		fmt.Println(index, " -  - ", newID)
	}

	// for file
	file, err := os.Open("mail-ids.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		newID := strings.Replace(scanner.Text(), "@gmail.com", "", 1)
		fmt.Println(newID)
	}

	if err := scanner.Err(); err != nil {
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
