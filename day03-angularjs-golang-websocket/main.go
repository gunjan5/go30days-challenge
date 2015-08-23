package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var connections map[*websocket.Conn]bool

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			return
		}
	}
}
func wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	connections[conn] = true
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			return
		}
		log.Println(string(msg))
		sendAll(msg)

	}
}
func main() {

	connections = make(map[*websocket.Conn]bool)
	// handle all requests by serving a file of the same name
	fileHandler := http.FileServer(http.Dir("web/"))
	http.Handle("/", fileHandler)
	http.HandleFunc("/ws", wsHandler)

	//log.Printf("Running on localhost:8080")
	addr := fmt.Sprintf("127.0.0.1:8080")
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
