package main

import (
  "fmt"
  "time"

)

func heyyall(msg chan string){
	time.Sleep(1000*time.Millisecond)
	msg<-"bloody potatos"
}

func printmsg(msg string){
	fmt.Println(msg)
	fmt.Println("Ok now you've got time for me")
}


func main() {
	messages := make(chan string)
	//buffered channel for the commented part to work
	//messages := make(chan string, 4)

	//go func() { messages<-"ayy yo, can you hear me?"}()
	//go func() { messages<-"I will erase the first message bwahahahaha"}()

	//sending multiple messages in single (non-buffered) channel will cause an error
	// messages<-"ayy yo, can you hear me?"
	// messages<-"ayy yo, can you hear me2?"
	// messages<-"ayy yo, can you hear me3?"

	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)





    go heyyall(messages)

	go printmsg(<-messages)

	var input string
	fmt.Scanln(&input)
	fmt.Println("ok dude! the show is over, go home and eat a potato")
}
