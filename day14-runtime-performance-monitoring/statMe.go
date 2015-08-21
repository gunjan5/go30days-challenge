package main

import (
	"flag"
	"fmt"
	_ "github.com/bmhatfield/go-runtime-metrics"
	"runtime"
	"time"
)

func sayMyName(name string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("Oh na na what's my name? ", name)
		time.Sleep(100 * time.Millisecond) //doing some fake work here... please wait ZzzZzzZzzz
	}

}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("You've got", runtime.NumCPU(), " CPU cores yo!")

	sayMyName("basic")

	go sayMyName("Gopher")

	flag.Parse()

	go sayMyName("Rabbit")

	go sayMyName("aaaaaaaaa")

	go sayMyName("bbbbbbbbbb")

	var input string
	fmt.Scanln(&input)
	fmt.Println("ok dude! the show is over, go home and eat a potato")
}
