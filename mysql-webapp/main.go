package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0p-s4ectet"))

var loggedIn = false
var connString = "root:deepak@/golang_webapp?charset=utf8"

type page struct {
	Title      string
	Comments   []string
	IsLoggedIn bool
	Errors     []string
	Msg        []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homeGet(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]

	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	}

	db, err := sql.Open("mysql", connString)
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT content FROM comment")
	checkError(err)

	var comments []string
	for rows.Next() {
		var content string
		err = rows.Scan(&content)
		checkError(err)
		comments = append(comments, content)
	}
	var errors []string
	if flashes := session.Flashes("danger"); len(flashes) > 0 {
		fmt.Println(flashes)
	}

	homepage := page{Title: "Home Page", Comments: comments, IsLoggedIn: loggedIn, Errors: errors}
	err = tpl.ExecuteTemplate(w, "index.html", homepage)
	checkError(err)
}

func homePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	db, err := sql.Open("mysql", connString)
	checkError(err)
	defer db.Close()

	stat, err := db.Exec("INSERT into comment set content = ?", comment)
	checkError(err)
	fmt.Println(stat.LastInsertId())

	http.Redirect(w, r, "/", 302)
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	homepage := page{Title: "Login Page", IsLoggedIn: loggedIn}
	err := tpl.ExecuteTemplate(w, "login.html", homepage)
	checkError(err)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	// check if username and password exists
	db, err := sql.Open("mysql", connString)
	checkError(err)
	defer db.Close()

	var dbUsername string
	var dbPassword string

	rows := db.QueryRow("SELECT username, password FROM user WHERE username = ?", username).Scan(&dbUsername, &dbPassword)
	fmt.Println(rows, dbUsername, dbPassword)
	if dbUsername == "" {
		session, _ := store.Get(r, "session")
		session.AddFlash("Invalid username", "danger")
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}

	hash, err := HashPassword(password)
	checkError(err)

	fmt.Println(dbPassword, hash)
	match := CheckPasswordHash(password, dbPassword)

	if match == true {
		fmt.Println(match)
		session, _ := store.Get(r, "session")
		session.Values["username"] = username
		session.Save(r, w)

		loggedIn = true
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]

	if !ok {
		return
	}

	username, ok := untyped.(string)

	if !ok {
		return
	}

	w.Write([]byte(username))
}

func registerGet(w http.ResponseWriter, r *http.Request) {
	homepage := page{Title: "Register Page", IsLoggedIn: loggedIn}
	err := tpl.ExecuteTemplate(w, "register.html", homepage)
	checkError(err)
}

func registerPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	hash, err := HashPassword(password)
	checkError(err)
	// register users
	db, err := sql.Open("mysql", connString)
	checkError(err)
	defer db.Close()

	var oldUsername string
	err = db.QueryRow("SELECT username FROM user where username = ?", username).Scan(&oldUsername)

	if oldUsername != "" {
		session, _ := store.Get(r, "session")
		session.AddFlash("username already exists", "danger")
		session.Save(r, w)
	} else {
		_, err := db.Exec("INSERT INTO user set username = ?, password = ?", username, hash)
		checkError(err)
	}

	http.Redirect(w, r, "/login", 302)
}

func logout(w http.ResponseWriter, r *http.Request) {
	loggedIn = false

	session, _ := store.Get(r, "session")
	session.Values["username"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", 302)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeGet).Methods("GET")
	r.HandleFunc("/", homePost).Methods("POST")

	r.HandleFunc("/login", loginGet).Methods("GET")
	r.HandleFunc("/login", loginPost).Methods("POST")

	r.HandleFunc("/logout", logout).Methods("GET")

	r.HandleFunc("/test", handlerTest).Methods("GET")

	r.HandleFunc("/register", registerGet).Methods("GET")
	r.HandleFunc("/register", registerPost).Methods("POST")

	fs := http.FileServer(http.Dir("./static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
