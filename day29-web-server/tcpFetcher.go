package main


import (
	"fmt"
	"net"
	"bufio"
	"io/ioutil"
)



func main() {
	conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
	panic(err)
}
fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')


_,err=ioutil.ReadAll(conn)
if err != nil {
	panic(err)
}


fmt.Println(status)
//fmt.Println(string(v))

}