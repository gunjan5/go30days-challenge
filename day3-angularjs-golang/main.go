package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// handle all requests by serving a file of the same name
	fileHandler := http.FileServer(http.Dir("web/"))
	http.Handle("/", fileHandler)

	log.Printf("Running on localhost:8080")
	addr := fmt.Sprintf("127.0.0.1:8080")
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
