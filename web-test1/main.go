package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	gu "github.com/deepaksinghkushwah/tests/web-test1/generalutility"
	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	ID        int
	Username  string
	Firstname string
	Lastname  string
	Email     string
}

func home(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	t, err := template.ParseFiles("form.html")
	checkError(err)

	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(template.HTMLEscapeString(r.Form["username"][0]))
	}

	db, err := sql.Open("mysql", "root:deepak@/test")
	checkError(err)

	/*stmt, err := db.Prepare("INSERT INTO persons set firstname = ?, lastname = ?, email = ?, `username` = ?, password = ?")
	checkError(err)

	for i := 1; i < 10000; i++ {
		inc := fmt.Sprintf("%v", i)
		_, err := stmt.Exec("Test", "User"+inc, "testuser"+inc+"@localhost.com", "testuser"+inc, "123456")
		checkError(err)
	}*/

	rows, err := db.Query("SELECT id, username,firstname, lastname, email from persons")
	checkError(err)

	p := make(map[int]person)
	counter := 0
	for rows.Next() {
		var id int
		var username string
		var firstname string
		var lastname string
		var email string
		err = rows.Scan(&id, &username, &firstname, &lastname, &email)
		checkError(err)
		p[counter] = person{ID: id, Username: username, Firstname: firstname, Lastname: lastname, Email: email}
		counter++
	}

	t.Execute(w, p)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "<hr/>Time Taken: "+fmt.Sprintf("%v", elapsed))
}

func main() {
	c := gu.GetPrice()
	fmt.Println(c)
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", nil)
	checkError(err)
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
