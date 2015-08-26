package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("meNew.txt")
	if err != nil {
		panic(err)
	}
	log.Println(newFile)
	newFile.Close()
}
