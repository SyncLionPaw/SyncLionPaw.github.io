# bufio 带缓冲的 IO 操作库

bufio是标准库的一个包，全称为 buffered I/O，主要用于提供带缓冲的 I/O 操作，能够提高 I/O 操作的效率，减少系统调用的次数。


## 带缓冲区

bufio.NewReader函数可以将一个 io.Reader 包装成一个 bufio.Reader。它会维护一个内部的缓冲区，当从 bufio.Reader 中读取数据时，它会先从缓冲区中读取，如果缓冲区中的数据不足，才会从底层的 io.Reader 中读取更多数据填充缓冲区。

bufio.NewWriter函数可以将一个 io.Writer 包装成一个 bufio.Writer。它同样维护一个内部的缓冲区，当向 bufio.Writer 写入数据时，数据会先写入缓冲区，当缓冲区满了或者调用了 Flush 方法时，才会将缓冲区中的数据写入到底层的 io.Writer 中

这个缓冲区域是维护在内存中，所以相较于多次的直接系统调用和磁盘、网络打交道，要相对快一些。

### 写缓冲区的存在

使用 bufio.NewWriter 证明：是先向缓冲区写，Flush之后，才向Stdout写
```go
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
```
Flush 方法用于将缓冲区中的数据全部写入到底层的 io.Writer。
当使用 bufio.NewWriter(stdout) 的时候，会默认创建一个4096大小的缓冲区。
并且他的参数os.Stdout 本身是一个io.Writer(具体类型是os.File)。

### 读缓冲区
bufio.Reader 也是维护了一个读缓冲区，先尽量从读缓冲区读。

```go
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
```
### bufio.Scanner

bufio.Scanner 是一个方便的工具，用于逐行读取文本。它也可以将一个 io.Reader 包装起来，通过设置不同的分隔符（默认是换行符）来读取数据。

```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    fmt.Println("Scanned line:", scanner.Text())
}
if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```
这里通过 bufio.NewScanner 包装了标准输入流 os.Stdin，然后通过循环调用 Scan 方法逐行读取用户输入，并通过 Text 方法获取每行的内容。

## 总结
bufio是Go的标准的包，维护缓冲区，用来避免额外的IO操作。
- 提高效率：通过缓冲机制减少了对底层 I/O 系统调用的次数，提高了 I/O 操作的效率。
- 方便使用：提供了丰富的接口和方法，方便开发者进行各种 I/O 操作，比如逐行读取、按分隔符读取等, 比bare的io.Reader 例如 strings.Reader 要方便一点。

bufio 包在处理文件 I/O、网络 I/O 等场景中非常有用，能够帮助开发者更高效地进行数据读取和写入操作。

## 疑问：
对于本身就在内存的数据，例如string, 何必还要"减少底层的strings.NewReader 的调用次数"？

- 对于内存中的数据（如 string），bufio.Reader 的缓冲机制并不会带来显著的性能提升，因为内存读取本身已经非常快。

- 如果你需要逐行读取或实现复杂的读取逻辑，bufio.Reader 或 bufio.Scanner 仍然是非常有用的工具，因为它们提供了更方便的方法和灵活性。