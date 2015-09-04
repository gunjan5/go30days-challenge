package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrorEmptyString = errors.New("Not gonna yell an epmpty message yo!, gimme something more meaningful") //exported type
)

func yellit(msg string) error {

	if msg == "" {
		return ErrorEmptyString
		// panic(ErrorEmptyString) //do this only if something is really bad panic and recover aint so good for normally bad things

	}
	_, err := fmt.Printf("%s\n", msg)

	return err

}

func main() {

	if err := yellit(""); err != nil {
		fmt.Printf("yelling didnt work dude: %s\n", err)
		os.Exit(1)
	}
}
