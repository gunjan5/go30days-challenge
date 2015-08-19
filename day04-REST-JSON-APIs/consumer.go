package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
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

func main() {
	url := "http://localhost:8080"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload


	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Stuff.Movie, " \n", p.Stuff.TVShow)
	
}