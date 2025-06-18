# embed

## 引入
有一次在项目中看到了这样的内容

```go
//go:embed jules_verne.txt
//go:embed vocab
```
看起来像是注释，但是是一种特殊的语言特性。

## 作用
embed 的主要用途是将外部文件的内容嵌入到 Go 程序中，使得程序可以直接使用这些文件内容，而无需在运行时加载外部文件。以下是一个简单的例子：

## 示例

示例：嵌入 HTML 文件
假设我们有一个 HTML 文件 index.html，内容如下：

```html
<!DOCTYPE html>
<html>
<head>
    <title>Embedded Example</title>
</head>
<body>
    <h1>Hello, Go Embed!</h1>
</body>
</html>
```

我们可以使用 embed 将这个文件嵌入到 Go 程序中：
```go
package main

import (
    _ "embed"
    "fmt"
)

//go:embed index.html
var htmlContent string

func main() {
    // 输出嵌入的 HTML 内容
    fmt.Println("Embedded HTML Content:")
    fmt.Println(htmlContent)
}
```

## 总结

写在谁的头上，那个变量就变成了这个嵌入文件的内容

用途
嵌入静态资源:

嵌入 HTML、CSS、JavaScript 文件，用于构建 Web 应用。
嵌入配置文件（如 JSON、YAML）。
嵌入文本数据:

嵌入词汇表、文档或其他文本数据（如你的 jules_verne.txt 和 vocab 文件）。
简化部署:

将所有资源打包到一个二进制文件中，避免运行时依赖外部文件。

注意事项
嵌入的文件路径必须是相对路径。
嵌入的文件内容会增加二进制文件的大小，适合小型文件。
embed 仅支持 Go 1.16 及以上版本。
通过 embed，你可以轻松地将外部资源整合到 Go 程序中，简化开发和部署流程。