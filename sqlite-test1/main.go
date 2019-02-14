package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./database.sqlite3")

	var id int
	var username, password, email string

	// insert into table
	/*stmt, _ := db.Prepare("INSERT INTO user (username, email, password) values (?, ?, ?)")
	for i := 0; i < 2000; i++ {
		stmt.Exec("sam"+strconv.Itoa(i), "sam"+strconv.Itoa(i)+"@localhost.com", "123456")
	}

	defer stmt.Close()*/

	rows, _ := db.Query("SELECT id, username, password, email FROM user WHERE id > ?", 0)
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
	stmt2, _ := db.Prepare("SELECT id, username, password, email FROM user where id = ?")
	stmt2.QueryRow("MAX(id)").Scan(&id, &username, &password, &email)
	defer stmt2.Close()
	fmt.Println("--------------------------------------")
	fmt.Println("Fetching single row where last id is", id)
	fmt.Println("--------------------------------------")
	fmt.Println(id, username, password, email)
}
