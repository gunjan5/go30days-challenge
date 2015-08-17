package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)
	mutex := new(sync.Mutex)

	for i := 1; i < 10; i++ {
		//fmt.Println(i)
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func() {
				fmt.Printf("%d %d\n", i, j)
				mutex.Unlock()

			}()
		}
	}

}

//Not sure why it prints the value 10 even with the use of Mutex
