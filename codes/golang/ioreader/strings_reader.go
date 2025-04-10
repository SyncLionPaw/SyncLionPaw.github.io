// strings.NewReader 创建的对象，实现了 io.Reader 接口。
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// io.Reader interface
	r := strings.NewReader("hello world!")

	buf := make([]byte, 5) // 缓冲区
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("read %v byte into buf, read content:%v\n", n, string(buf[:n]))
		}
		if err == io.EOF {
			fmt.Println("reached end of file.")
			break
		}
		if err != nil {
			fmt.Println("error encounter: %v", err.Error())
			break
		}
	}
}

/*
read 5 byte into buf, read content:hello
read 5 byte into buf, read content: worl
read 2 byte into buf, read content:d!
reached end of file.
*/
