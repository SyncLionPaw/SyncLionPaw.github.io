package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout) // 能向标准输出写
	w.WriteString("this is a string.")
	w.WriteByte(97)
	smile := '😀'

	fmt.Printf("unicode: %U\n", smile)

	w.WriteRune(smile) // 这个是一个 unicode 码点

	fmt.Println("** up to now, writer content hasn't been printed to stdout **")

	w.Flush()
}
