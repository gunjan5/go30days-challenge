// http://www.json.org/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Anything interface{}

func main() {
	jsonData := `
	{
	"ans_to_life" : 42,
	"best_pokemon" : "charizard"
	}
	 `
	var obj map[string]Anything
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println("answer to all questions is: ", obj["ans_to_life"])

	v, ok := obj["ans_to_life"].(float64) //v, ok := obj["best_pokemon"].(float64) //asserting the type to float64
	if !ok {
		v = 0 //take default value instead of panicing & log error
		log.Println("ay man, we've got an issue with asserting the data type, but I'll work something out instead of panicing :) ")
	}
	x := 100 + v

	fmt.Printf("Type of ans_to_life: %T, value of x: %v\n", v, x)

}
