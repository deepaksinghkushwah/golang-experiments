package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.sqlite3")
	checkError(err)
	var id int
	var username, password, email string

	// insert into table
	stmt, err := db.Prepare("INSERT INTO user (firstname,lastname,username, email, password) values (?, ?, ?, ?, ?)")
	checkError(err)
	for i := 20; i <= 30; i++ {
		_, err := stmt.Exec("sam", strconv.Itoa(i), "sam"+strconv.Itoa(i), "sam"+strconv.Itoa(i)+"@localhost.com", "123456")
		checkError(err)
	}

	defer stmt.Close()

	rows, err := db.Query("SELECT id, username, password, email FROM user WHERE id > ?", 0)
	checkError(err)
	defer rows.Close()
	// fetching multiple rows
	for rows.Next() {

		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(id, username, password, email)
	}

	// fetching single row and display its content
	stmt2, err := db.Prepare("SELECT id, username, password, email FROM user where id = ?")
	checkError(err)
	stmt2.QueryRow("MAX(id)").Scan(&id, &username, &password, &email)
	defer stmt2.Close()
	fmt.Println("--------------------------------------")
	fmt.Println("Fetching single row where last id is", id)
	fmt.Println("--------------------------------------")
	fmt.Println(id, username, password, email)
}

func checkError(e error) {
	if e != nil {
		log.Println(e)
	}
}
