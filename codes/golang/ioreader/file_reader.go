// file 实现了 io.Reader 接口
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	buf := make([]byte, 128)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Printf("read %v bytes: %v\n", n, string(buf[:n]))
		}
		if err == io.EOF {
			fmt.Println("reached end of file")
			break
		}
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
			break
		}
	}
}
