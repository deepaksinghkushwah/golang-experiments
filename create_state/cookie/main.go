package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}
	counter, _ := strconv.Atoi(c.Value)
	counter++
	c.Value = strconv.Itoa(counter)
	http.SetCookie(w, c)

	io.WriteString(w, "Cookie counter: "+string(c.Value))
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}
