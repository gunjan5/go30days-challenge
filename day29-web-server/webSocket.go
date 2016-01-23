package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
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

	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {

			for {
				_, msg, _ := conn.ReadMessage()
				fmt.Println(string(msg))
			}

		}(conn)

	})

		http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			ch:= time.Tick(5* time.Second)
			for range ch{
				conn.WriteJSON(myStruct{
					User: "gunjan5",
					Name: "Gunjan Patel",
					})
			}


		}(conn)

	})

	http.ListenAndServe(":8080", nil)
}
type myStruct struct{
	User string `json:"user"`
	Name string `json:"name"`
}
