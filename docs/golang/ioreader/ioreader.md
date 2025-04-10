# io.Reader
写go的时候经常见到 io.Reader 这个东西，那么这个是什么呢？

## 是一个接口
io.Reader 是一个接口 interface 在 go 的 io 包里面定义的。

```go
// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// If len(p) == 0, Read should always return n == 0. It may return a
// non-nil error if some error condition is known, such as EOF.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

## 接口只定义了一个Read

只需要实现了 Read 方法的就是 Reader。


## 谁实现了这个接口
有很多常用的工具，都实现了这个借口。

### strings.NewReader
strings.NewReader 创建的对象，就是实现了Reader接口的。
```go
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
```

### os.File
os.File 实现了 io.Reader 接口
```go
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

```
### net.Conn 实现了 io.Reader
可以从网络连接中读取字节流。
```go
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
```

## 总结 io.Reader
io.Reader 是一个接口，只需要实现一个 Read 方法，就实现了这个接口。
很多库都实现了这个接口，例如
- strings.Reader 用来读 string
- net.Conn 用来读网络连接
- os.File 用来读文件

不同的实际类型，可以各自实现Read，只是实现方式不一样。

