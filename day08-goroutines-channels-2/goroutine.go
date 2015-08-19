package main



import (
   "fmt"
   "time"
   "runtime"

)

func doStuff(){
	msg := make(chan string)

	go func(){
		println("yello")
	}()

	go func(){
		fmt.Println("#####")
		time.Sleep(0*time.Millisecond)
		msg<-"............."
	}()

	//time.Sleep(100*time.Millisecond)

	println(<-msg)
}

func main() {

	runtime.GOMAXPROCS(2)

	for i:=0; i<100; i++ {
		doStuff()
	}
	


}