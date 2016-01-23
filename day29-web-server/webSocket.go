package main

import (
	_ "fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		 conn, _ := upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {

			for {
				mType, msg, _ := conn.ReadMessage()
				conn.WriteMessage(mType, msg)
			}

		}(conn)

	})

	http.ListenAndServe(":8080", nil)
}
