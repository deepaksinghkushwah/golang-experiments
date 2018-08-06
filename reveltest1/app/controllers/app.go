package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	type person struct {
		ID        int
		Firstname string
	}
	var persons = make(map[int]person)
	db, err := sql.Open("mysql", "root:deepak@/test")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("Select id, firstname from persons")
	if err != nil {
		log.Println(err)
	}
	counter := 0
	for rows.Next() {
		var id int
		var firstname string
		err = rows.Scan(&id, &firstname)
		if err != nil {
			log.Println(err)
		}
		persons[counter] = person{ID: id, Firstname: firstname}
		counter++
	}

	return c.Render(persons)
}
