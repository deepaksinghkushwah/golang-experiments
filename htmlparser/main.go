package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	chUrls := make(chan string)
	chFinidhed := make(chan bool)

	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinidhed)
	}

	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinidhed:
			c++
		}
	}

	fmt.Println("\nFound", len(foundUrls), "unique urls:\n ")

	for url := range foundUrls {
		fmt.Println(" - " + url)
	}

	close(chUrls)

}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func crawl(url string, ch chan string, chFinidhed chan bool) {
	resp, err := http.Get(url)
	defer func() {
		chFinidhed <- true
	}()

	checkError(err)
	b := resp.Body
	defer b.Close()

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data != "a" {
				continue
			}

			ok, url := getHref(t)
			if !ok {
				continue
			}

			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}
		}
	}
}

func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}
