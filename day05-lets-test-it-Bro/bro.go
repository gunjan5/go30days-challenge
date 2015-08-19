package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

// Just return folder where `bro` has running
func GetCurrentLocation() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

type Bro struct {
	GoExt    string
	Location string
	Files    map[string]time.Time
}

func (b *Bro) WalkLocation() {

	Recurse := true
	b.Files = make(map[string]time.Time)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}

		if info.IsDir() && path != b.Location && !Recurse {
			return filepath.SkipDir
		}

		if filepath.Ext(path) == b.GoExt {
			b.Files[path] = info.ModTime()
		}
		return nil
	}

	filepath.Walk(b.Location, walkFunc)
	log.Printf("[INFO]: Eye on %d files\n", len(b.Files))
}

func (b *Bro) HelpMe() {
	for {
		b.LookThem()

		time.Sleep(time.Duration(2 * time.Second))
	}
}

func (b *Bro) LookThem() {
	for file, modtime := range b.Files {
		stat, err := os.Stat(file)
		if err != nil {
			log.Fatal("[ERROR]: ", err.Error())
		}

		ntime := stat.ModTime()
		if ntime.Sub(modtime) > 0 {
			b.Files[file] = ntime
			log.Printf("[INFO]: Changed files %s\n", file)
			PrepareCmd(file)
		}

	}
}

func IsTestFile(nameFile string) bool {
	matched, _ := regexp.MatchString("._test.go", nameFile)
	return matched
}

func HasTestFile(dirFile, nameFile string) bool {
	testFilePattern := "%s_test.go"
	fullPath := filepath.Join(dirFile, fmt.Sprintf(testFilePattern, nameFile))

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func PrepareCmd(path string) bool {
	nameFile := filepath.Base(path)
	dirFile := filepath.Dir(path)

	var mainFile string
	var testFile string

	if IsTestFile(nameFile) == true {
		testFile = path
		mainFile = filepath.Join(dirFile, strings.Replace(nameFile, "_test.go", ".go", -1))
		return TestCommand(mainFile, testFile)
	}

	if HasTestFile(dirFile, strings.Replace(nameFile, ".go", "", -1)) == true {
		testFile = filepath.Join(dirFile, strings.Replace(nameFile, ".go", "_test.go", -1))
		mainFile = path
		return TestCommand(mainFile, testFile)
	}

	log.Printf("[WARNING]: %s need of test man!", nameFile)
	return false

}

func TestCommand(mainFile, testFile string) bool {

	cmd := exec.Command("go", "test", mainFile, testFile, "-v")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "bro"
	app.Usage = "I'll help you with tests"
	app.Action = func(c *cli.Context) {
		var where string
		if c.Args().Present() == false {
			where = GetCurrentLocation()
		} else {
			where = c.Args().First()
		}

		bro := Bro{GoExt: ".go", Location: where}
		bro.WalkLocation()

		bro.HelpMe()
	}

	app.Run(os.Args)
}
