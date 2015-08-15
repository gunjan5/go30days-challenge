package main

import (
  "fmt"

)




func main() {
	messages := make(chan string)

	go func() { messages<-"ayy yo, can you hear me?"}()

	//go func() { messages<-"I will erase the first message bwahahahaha"}()


	msg := <-messages
	fmt.Println(msg)
}
