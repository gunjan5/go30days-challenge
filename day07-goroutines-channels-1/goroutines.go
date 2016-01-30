package main

import (
  "fmt"
  "runtime"
  "time"
)

func init() {
	
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func sayMyName(name string){
	for i:=0; i<5; i++ {
		runtime.Gosched() //Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
		fmt.Println("Oh na na what's my name? ", name)
		time.Sleep(100*time.Millisecond) //doing some fake work here... please wait ZzzZzzZzzz
	}
	
}

func main() {


	fmt.Println("You've got", runtime.NumCPU() , " CPU cores yo!")

	sayMyName("basic")

	go sayMyName("Gopher")

	go sayMyName("Rabbit")

	go sayMyName("aaaaaaaaa")

	go sayMyName("bbbbbbbbbb")

	// var input string
	// fmt.Scanln(&input)
	fmt.Println("ok dude! the show is over, go home and eat a potato")
}
