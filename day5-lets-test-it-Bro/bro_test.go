package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func createTestFolder(withTestFile bool) {
	testFolder := "test_folder"
	err := os.Mkdir(testFolder, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(fmt.Sprintf("%s/%s", testFolder, "other_folder"), 0777)
	if err != nil {
		log.Fatal(err)
	}

	// creating files
	_, err = os.Create(fmt.Sprintf("%s/%s", testFolder, "fileone.go"))
	if err != nil {
		log.Fatal(err)
	}

	if withTestFile == true {
		_, err = os.Create(fmt.Sprintf("%s/%s", testFolder, "fileone_test.go"))
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Create(fmt.Sprintf("%s/%s/%s", testFolder, "other_folder", "filetwo.go"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create(fmt.Sprintf("%s/%s", testFolder, "filethree.go"))
	if err != nil {
		log.Fatal(err)
	}

}

func dropTestFolder() {
	err := os.RemoveAll("test_folder")
	if err != nil {
		log.Fatal(err)
	}
}

func TestWalkLocation(t *testing.T) {
	withTestFiles := false
	createTestFolder(withTestFiles)

	here, _ := filepath.Abs(".")
	here_test := filepath.Join(here, "test_folder")

	bro := Bro{GoExt: ".go", Location: here_test}
	bro.WalkLocation()

	expected := []string{
		fmt.Sprintf("%s/%s", here_test, "fileone.go"),
		fmt.Sprintf("%s/%s", here_test, "filethree.go"),
		fmt.Sprintf("%s/%s", here_test, "other_folder/filetwo.go"),
	}

	for _, file := range expected {
		if _, ok := bro.Files[file]; !ok {
			t.Errorf("expected %s", file)
		}
	}
	dropTestFolder()
}

func TestIsTestFile(t *testing.T) {
	withTestFiles := true
	createTestFolder(withTestFiles)

	here, _ := filepath.Abs(".")
	here_test := filepath.Join(here, "test_folder")

	normal_file := IsTestFile(fmt.Sprintf("%s/%s", here_test, "fileone.go"))
	test_file := IsTestFile(fmt.Sprintf("%s/%s", here_test, "fileone_test.go"))

	if normal_file == true {
		t.Error("Expected False result")
	}
	if test_file == false {
		t.Error("Expected True result")
	}
	dropTestFolder()
}

func TestHasTestFile(t *testing.T) {
	withTestFiles := true
	createTestFolder(withTestFiles)

	here, _ := filepath.Abs(".")
	here_test := filepath.Join(here, "test_folder")

	hasTestFile := HasTestFile(here_test, "fileone")
	not := HasTestFile(fmt.Sprintf("%s/%s", here_test, "other_folder"), "filetwo")

	if hasTestFile == false {
		t.Error("Expected true result")
	}
	if not == true {
		t.Error("Expected false result")
	}
	dropTestFolder()
}
