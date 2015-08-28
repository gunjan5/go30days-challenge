package main

import (
	"log"
	"os"

	"github.com/jasonmoo/cryp"
)

func main() {

	key, exists := os.LookupEnv("CRYP_KEY")
	if !exists {
		log.Fatal("CRYP_KEY not set in environment")
	}

	for _, dir := range os.Args[1:] {
		if err := cryp.EncryptDirFiles(dir, []byte(key)); err != nil {
			log.Fatal(err)
		}
	}

}
