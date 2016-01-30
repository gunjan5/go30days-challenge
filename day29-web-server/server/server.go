package main

import (
	"fmt"
	"net/http"
	_"io"
	"os"
	"flag"
	"strconv"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello world!")
// }

// func testhandler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "TEST!!!!!!!!!!!!!!!!!!!!!!! !")
// }

func main() {
	//var p int


	mux:= http.NewServeMux()
	wd,_:=os.Getwd()
	dir := flag.String("document_root", wd, "document root")
	port := flag.Int("port", 8080, "specify server listen port")
	flag.Parse()
	
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
  fmt.Println("ERROR: File/path does not exist! Exiting the program ...")
  os.Exit(1)
}


	fileRoot:=http.Dir(*dir)
	//redirectHandler:= http.RedirectHandler(":8080/index.html", 307)
	files:=http.FileServer(fileRoot)
	//http.HandleFunc("/test", testhandler)
	mux.Handle("/", files)
	//fmt.Println(*port)
	srvAddr := ":"+strconv.Itoa(*port)
	//fmt.Printf("%T\n\n\n", srvAddr)
	fmt.Println(srvAddr)
	fmt.Println(fileRoot)
		
	http.ListenAndServe(srvAddr, mux)
}









// func main() {
// 	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/Volumes/Other/desktop/go30days-challenge/day29-web-server/"))))

// }







// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"strings"
// )

// func handleConn(conn net.Conn) {
// 	//defer conn.Close()
// 	scanner := bufio.NewScanner(conn)
// 	i := 0
// 	for scanner.Scan() {
// 		ln := scanner.Text()
// 		fmt.Println(ln)

// 		if i == 0 {
// 			method := strings.Fields(ln)[0]
// 			fmt.Println("METHOD", method)
// 		} else {
// 			// in headers now
// 			// when line is empty, header is done
// 			if ln == "" {
// 				break
// 			}
// 		}

// 		i++
// 	}

// 	// response
// 	body := "hello world 2"

// 	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
// 	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
// 	io.WriteString(conn, "\r\n")
// 	io.WriteString(conn, body)
// }

// func main() {
// 	server, err := net.Listen("tcp", ":9000")
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	defer server.Close()

// 	for {
// 		conn, err := server.Accept()
// 		if err != nil {
// 			log.Fatalln(err.Error())
// 		}
// 		go handleConn(conn)
// 	}
// }
