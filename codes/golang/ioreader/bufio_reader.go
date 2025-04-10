package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("bufio_test.txt") // reader
	br := bufio.NewReader(r)          // bufio 包装 io.reader

	shortStr, _ := br.ReadString('G')
	fmt.Printf("shortStr: %v\n", shortStr)

	str, _, _ := br.ReadLine()
	fmt.Printf("str: %v\n", str)

	// Buffered, 返回已经预读缓存的字节数
	// Buffered returns the number of bytes that can be read from the current buffer
	fmt.Printf("br.Buffered(): %v\n", br.Buffered()) // 2804
}
