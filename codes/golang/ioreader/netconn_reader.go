// net.Conn 实现了 io.Reader 接口

package main

import (
	"fmt"
	"io"
	"net"
)


func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	fmt.Println("server is listening on port 8080...")
	conn, _ := listener.Accept()

	defer conn.Close()

	buf := make([]byte, 128)

	for {
		n, err := conn.Read(buf)
		if n > 0 {
			fmt.Printf("read %v bytes: %s\n", n, string(buf[:n]))
		}
		if err == io.EOF {
			fmt.Println("client closed the connection")
		}
		if err != nil {
			fmt.Printf("error reading from connection %v\n", err.Error())
		}
	}
}