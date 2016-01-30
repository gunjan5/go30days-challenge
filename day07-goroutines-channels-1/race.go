package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
	"math/rand"
)



func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}



var globalcounter int
var	wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incr("Red")
	go incr("Blue")
	wg.Wait()
	fmt.Println("Final counter:", globalcounter, "instead of 100")
}

func incr(s string) {
	for i := 0; i < 50; i++ {
		x:=globalcounter
		x++
		time.Sleep(rand.Intn(2)*time.Millisecond) //OR runtime.Gosched() //Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
		globalcounter = x
		fmt.Println(s, i, "counter:", globalcounter)
	}
	wg.Done()
}


