package main


import (
	"fmt"
	"net"
	"bufio"
	"io/ioutil"
)



func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	panic(err)
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')


v,err:=ioutil.ReadAll(conn)
if err != nil {
	panic(err)
}


fmt.Println(status)
fmt.Println(string(v))

}