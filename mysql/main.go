package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:deepak@/test?charset=utf8")
	checkError(err)
	defer db.Close()

	//stmt, err := db.Prepare("SELECT * FROM person")
	//checkError(err)

	rows, err := db.Query("SELECT username FROM persons")
	checkError(err)

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		checkError(err)
		fmt.Println(username)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
