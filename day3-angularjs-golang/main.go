package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

func wsHandler ( w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w,r, nil, 1024, 1024)
	if _,ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
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
	addr := fmt.Sprintf("127.0.0.1:80")
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
