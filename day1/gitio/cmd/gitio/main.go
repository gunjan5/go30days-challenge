package main

import (
	"fmt"
	"os"

	"github.com/xlab/gitio"
	"gopkg.in/mflag.v1"
)

var (
	longURL  string
	codeOpt  string
	forceOpt bool
)

func init() {
	mflag.StringVar(&codeOpt, []string{"c", "-code"}, "", "A custom code for the short link, e.g. http://git.io/mycode")
	mflag.BoolVar(&forceOpt, []string{"f", "-force"}, false, "Try to shorten link even if the custom code has been used previously.")

	mflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <long url>\n", os.Args[0])
		mflag.PrintDefaults()
	}
	mflag.Parse()
	if longURL = mflag.Arg(0); len(longURL) == 0 {
		mflag.Usage()
		return
	}
}

func main() {
	if len(codeOpt) > 0 && !forceOpt {
		taken, err := gitio.CheckTaken(codeOpt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		if taken {
			fmt.Printf("warning: custom code \"%s\" has already been taken.\n", codeOpt)
			fmt.Println("tip: use the --force flag to suppress this check.")
			os.Exit(1)
		}
	}
	shortURL, err := gitio.Shorten(longURL, codeOpt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(shortURL)
}
