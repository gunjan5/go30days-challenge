package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gunjan5/go30days-challenge/day1/gitio"
)

var (
	longUrl  string
	)

type Page struct {
    		Title string
    		Body  []byte
    	}

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    longUrl = r.URL.Path[1:]
    fmt.Println(longUrl)
    //longUrl = "https://github.com/gunjan5/go30days-challenge/tree/master/day1"
    shortUrl, err := gitio.Shorten(longUrl)
    if err != nil {
    	fmt.Fprintf(w, "Oh man! You broke it! %s", err)
    } else {
    	fmt.Fprintf(w, "%s", shortUrl)
    }
    

}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":8080", nil)
    // longUrl = "https://github.com/gunjan5/go30days-challenge/tree/master/day1"
    // shortUrl, err := gitio.Shorten(longUrl)
    // if err != nil {
    // 	fmt.Printf("Oh man! You broke it! %s", err)
    // }
    // fmt.Println(shortUrl)
}
