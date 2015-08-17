package main

import (
	"fmt"
)

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

func main() {

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"frodo@shiar.io"},
		From:    "gandalf@whitecouncil.org",
		Content: "You must keep it a sectet",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Interrupted by the black riders",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No message received")
	}
}
