package app

import (
	"log"
	"net/http"

	"github.com/deepaksinghkushwah/site1/models"
	_ "github.com/go-sql-driver/mysql"
)

// ListAllBlogs will display all blogs
func ListAllBlogs(w http.ResponseWriter, r *http.Request) {

	dataToPass := []models.BlogList{}

	db := models.GetDbo()
	defer db.Close()

	err := db.Ping()
	models.CheckDbErr(err)

	rows, err := db.Query("SELECT * FROM blog")
	defer rows.Close()
	models.CheckDbErr(err)

	var id int32
	var title string
	var content string
	var status int32
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&id, &title, &content, &status)
			models.CheckDbErr(err)
			dataToPass = append(dataToPass, models.BlogList{ID: id, Title: title, Content: content, Status: status})
			//fmt.Println(title)
		}
	}

	err = rows.Err()
	models.CheckDbErr(err)

	page := models.BlogPageList{Title: "Blog", Keywords: "Blog Listing", Description: "Listing of blogs", Rows: dataToPass}
	err2 := tpl.ExecuteTemplate(w, "blog-list.html", page)
	if err2 != nil {
		log.Println(err2)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
