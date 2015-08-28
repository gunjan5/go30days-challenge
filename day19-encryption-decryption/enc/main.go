package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"

	"github.com/jasonmoo/cryp"
)

func main() {

	key, exists := os.LookupEnv("CRYP_KEY")
	if !exists {
		log.Fatal("CRYP_KEY not set in environment")
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	output, err := cryp.Encrypt(data, []byte(key))
	if err != nil {
		log.Fatal(err)
	}

	enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer enc.Close()
	enc.Write(output)

}
