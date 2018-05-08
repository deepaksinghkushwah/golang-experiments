package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	type blogList struct {
		ID         int32
		Title      string
		Content    string
		BlogStatus int32
	}
	dataToPass := []blogList{}
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM blog")

	checkErr(err)
	var id int32
	var title string
	var content string
	var status int32
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&id, &title, &content, &status)
			checkErr(err)
			dataToPass = append(dataToPass, blogList{ID: id, Title: title, Content: content, BlogStatus: status})
		}
	}
	rows.Close()
	db.Close()
	fmt.Println(dataToPass)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
