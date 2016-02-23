//See the Zipf's curve in action

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func main() {
	zipf := make([]int, 0)

	words := make(map[string]int)
	resp, _ := http.Get("https://blog.golang.org")

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
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
	for k, v := range words {
		if v > 50 {
			fmt.Println(k, " : ", v)
			zipf=append(zipf, v)

		}
	}

	sort.Ints(zipf)

	for _, v := range zipf {
		fmt.Println(v)
	}

}
