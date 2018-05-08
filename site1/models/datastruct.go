package models

import (
	"database/sql"
	"fmt"
)

// Page struct for creatng structured page
type Page struct {
	Title       string
	Keywords    string
	Description string
}

type BlogList struct {
	ID      int32
	Title   string
	Content string
	Status  int32
}

// BlogPageList for listing blog page
type BlogPageList struct {
	Title       string
	Keywords    string
	Description string
	Rows        []BlogList
}

// GetDbo function return database instance
func GetDbo() *sql.DB {
	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		CheckDbErr(err)
	}
	return db
}

// CheckDbErr is function which show error
func CheckDbErr(err error) {
	fmt.Println(err)
}
