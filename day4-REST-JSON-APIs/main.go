package main

import (
"encoding/json"
"fmt"
"net/http"
)


func serveRest ( w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", servRest)

	http.ListenAndServe("localhost:8080", nil)
	

}


