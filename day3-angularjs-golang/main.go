package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wsHandler ( w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w,r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage() 
		if err != nil {
			return
		}
		log.Println(string(msg))
		if err = conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}

	}
}
func main() {
	// handle all requests by serving a file of the same name
	fileHandler := http.FileServer(http.Dir("web/"))
	http.Handle("/", fileHandler)
	http.HandleFunc("/ws", wsHandler)

	log.Printf("Running on localhost:8080")
	addr := fmt.Sprintf("127.0.0.1:8080")
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
