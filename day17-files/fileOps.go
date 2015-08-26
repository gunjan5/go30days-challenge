package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

//use array [...]
//slices
//performance

type Gfile struct {
	Name    string
	Size    float32
	Mode    os.FileMode
	Content []string

	//permissions, owners, last modified
}

func (f *Gfile) createFile() {
	newFile, err = os.Create(f.Name)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	log.Println(newFile)

}

func (f *Gfile) changemod(mode os.FileMode) {
	err := os.Chmod(f.Name, mode)
	if err != nil {
		panic(err)
	}
	f.Mode = mode //should this be in a defer?
	log.Println("changed the file mode to", f.Mode, "for file: ", f.Name)

}

func main() {
	f1 := Gfile{Name: "newGfile.txt", Content: []string{"line1", "line2", "line3"}}
	f1.createFile()
	f1.changemod(0777)

}
