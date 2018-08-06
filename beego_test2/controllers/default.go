package controllers

import (
	"database/sql"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type person struct {
	ID        int
	Firstname string
}

func (c *MainController) Mydata() {
	users := make(map[int]person)
	db, err := sql.Open("mysql", "root:deepak@/test")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := db.Query("SELECT id, firstname from t1")
	if err != nil {
		log.Fatalln(err)
	}
	counter := 0
	for rows.Next() {
		var id int
		var firstname string
		err = rows.Scan(&id, &firstname)
		if err != nil {
			log.Fatalln(err)
		}
		users[counter] = person{ID: id, Firstname: firstname}
		counter++
	}
	c.Data["users"] = users
	c.TplName = "mydata.tpl"
}
