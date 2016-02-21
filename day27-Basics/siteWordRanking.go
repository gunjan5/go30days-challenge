package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	words := make(map[string]int)
	resp, _ := http.Get("https://blog.golang.org")

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(respBody))
	r := bytes.NewReader(respBody)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		wScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		wScanner.Split(bufio.ScanWords)

		for wScanner.Scan() {
			_, ok := words[wScanner.Text()]
			if !ok { //doesnt exist in the map
				words[wScanner.Text()] = 1
			} else {
				words[wScanner.Text()]++
			}

		}
	}
for k,v := range words{
	fmt.Println(k, " : ", v)
}

}
