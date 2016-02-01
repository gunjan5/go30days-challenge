// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// }

// func main() {

// 	mutex := new(sync.Mutex)

// 	for i := 0; i < 10; i++ {
// 		//fmt.Println(i)
// 		for j := 0; j < 10; j++ {
// 			mutex.Lock()
// 			go func() {
// 				fmt.Printf("%d %d\n", i, j)
// 				mutex.Unlock()

// 			}()
// 		}
// 	}

// }

// //Not sure why it prints the value 10 even with the use of Mutex



package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			go func(m, n int) {
				fmt.Printf("%d %d\n", m, n)
			}(i, j)
		}
	}
	var a string
	fmt.Scanln(&a)
}
