package main

import (
    "fmt"
    "net/http"
    "github.com/gunjan5/go30days-challenge/day1/gitio"
)

var (
	longUrl  string
	)

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    //longUrl = r.URL.Path[1:]
    longUrl = "https://github.com/gunjan5/go30days-challenge/tree/master/day1"
    shortUrl, err := gitio.Shorten(longUrl)
    if err != nil {
    	fmt.Fprintf("Oh man! You broke it! %s", err)
    }
    fmt.Fprintf(shortUrl)

}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
