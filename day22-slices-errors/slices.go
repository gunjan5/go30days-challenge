package main

import (
	_ "bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open("test.txt")
	check(err)
	defer f.Close()

	b := make([]byte, 100)
	n, err := f.Read(b)

	stringified := string(b)

	//fmt.Println(string(b))
	fmt.Printf("%d: % s \n", n, stringified)

	//now lets write

	//w, err := os.Open("writeMe.txt")
	// if err != nil {
	// 	fmt.Printf(" %s\n", err)
	// 	os.Exit(1)
	// }
	//defer w.Close()
	blob := "hi this is golang program slices.go, I'm trying to write to you Miss writeMe.txt, hope you like it!\n\n"

	err = ioutil.WriteFile("writeMe.txt", []byte(blob), 0777)
	check(err)

	//w.Write([]byte(blob))
}
