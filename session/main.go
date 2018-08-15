package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func home(w http.ResponseWriter, r *http.Request) {
	uid, _ := uuid.NewV4()
	cookie, err := r.Cookie("session-id")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    uid.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)

}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}
