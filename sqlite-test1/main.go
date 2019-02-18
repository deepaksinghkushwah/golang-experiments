package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}
	stmt, err := db.Prepare("INSERT INTO `user` (firstname, lastname, username, email, pass) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec("Deepak", "Kushwah", "deepak", "deepaksinghkushwah@gmail.com", "123456")
	if err != nil {
		log.Fatalln(err)
	}
	lid, _ := res.LastInsertId()
	fmt.Println(lid)
}
