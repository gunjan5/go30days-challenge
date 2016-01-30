
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
	//"math/rand"
	//"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 10000; i++ {
		runtime.Gosched() //Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
		//Alternative to Gosched() is //time.Sleep(time.Duration(rand.Intn(2))*time.Millisecond)

		atomic.AddInt64(&count, 1)
		//count++
		
		fmt.Println("Process: "+s+" printing:", i)
	}
	wg.Done()
}