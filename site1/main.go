package main

import (
	"fmt"
	"net/http"

	"github.com/deepaksinghkushwah/site1/app"
	"github.com/deepaksinghkushwah/site1/models"
)

func main() {

	//populateBlog()

	//fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/", app.Home)
	http.HandleFunc("/about", app.About)
	http.HandleFunc("/contact", app.Contact)
	http.HandleFunc("/bloglist", app.ListAllBlogs)
	http.Handle("favicon.ico", http.NotFoundHandler())

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func populateBlog() {
	db := models.GetDbo()
	for i := 0; i <= 999999; i++ {
		stmt, err := db.Prepare("INSERT INTO blog set title=?, content=?")
		if err != nil {
			models.CheckDbErr(err)
		}
		title := fmt.Sprintf("%s %d", "Title", i)
		content := fmt.Sprintf("%s %d", "Content", i)
		res, err := stmt.Exec(title, content)
		if err != nil {
			models.CheckDbErr(err)
		}
		fmt.Println(res.RowsAffected())
	}
}
