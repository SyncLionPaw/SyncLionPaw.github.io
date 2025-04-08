在golang函数的内部，可以定义另外的函数

```go
package main

import "fmt"

func main() {
	fmt.Println("Outer function")
	var innerFunction = func() {
        fmt.Println("Inner function")
    }
	innerFunction() // 没有这一句会出错，因为innerFunction定义而不使用
}
```

```go
package main

import "fmt"

func main() {
	fmt.Println("Outer function")
	innerFunction := func() { // 使用 := 也可以
        fmt.Println("Inner function")
    }
	innerFunction()
}
```