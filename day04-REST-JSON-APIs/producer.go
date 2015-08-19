package main

import (
"encoding/json"
"fmt"
"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Movie Movies
	TVShow Shows
}

type Movies map[string]float32
type Shows map[string]float32


func serveRest ( w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))


}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:8080", nil)

}

func getJsonResponse() ([]byte, error) {
	movies := make(map[string]float32)
	movies["Star Wars III"] = 9.3
	movies["Ex_Machina"] = 9.1
	movies["Inception"] = 8.7

	shows := make(map[string]float32)
	shows["Game of Thrones"] = 9.9
	shows["Silicon Valley"] = 10
	shows["the 100"] = 9.2

	d := Data{movies, shows}
	p := Payload{d}

	return json.MarshalIndent(p, "", " ")

}

