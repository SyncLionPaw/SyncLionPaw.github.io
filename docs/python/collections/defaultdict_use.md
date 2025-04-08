
Python 的 `defaultdict` 是 `collections` 模块中的一个非常有用的数据结构，它是字典（`dict`）的一个子类。`defaultdict` 的特殊之处在于，当你尝试访问一个不存在的键时，它不会抛出 `KeyError`，而是会使用一个默认值，这个默认值由你提供的一个工厂函数生成。

以下是 `defaultdict` 的基本用法：

### 导入 `defaultdict`

```python
from collections import defaultdict
```

### 创建 `defaultdict`

创建 `defaultdict` 时，你需要提供一个工厂函数，这个函数会在字典中需要一个不存在的键时被调用。

```python
# 使用 list 作为默认工厂函数
dd = defaultdict(list)
```

### 使用 `defaultdict`

```python
# 添加元素
dd['key1'].append(1)  # 因为 'key1' 不存在，所以首先会创建 'key1' 并用 [] 作为默认值
dd['key1'].append(2)

# 访问元素
print(dd['key1'])  # 输出: [1, 2]

# 尝试访问不存在的键
print(dd['key2'])  # 输出: []，因为 'key2' 不存在，所以使用默认的 [] 作为值
```

### 常用的默认工厂函数

- `list`: 当你想要一个列表作为默认值时。
- `set`: 当你想要一个集合作为默认值时。
- `int`: 当你想要一个整数作为默认值时，通常设置为 `int()`，它返回0。
- `str`: 当你想要一个字符串作为默认值时，通常设置为 `str()`，它返回空字符串`''`。
- `dict`: 当你想要一个字典作为默认值时。

### 示例：使用 `set` 作为默认工厂函数

```python
from collections import defaultdict

# 使用 set 作为默认工厂函数
dd = defaultdict(set)

# 添加元素
dd['key1'].add('value1')
dd['key2'].add('value2')

# 访问元素
print(dd['key1'])  # 输出: {'value1'}
print(dd['key2'])  # 输出: {'value2'}

# 尝试访问不存在的键
print(dd['key3'])  # 输出: set()，因为 'key3' 不存在，所以使用默认的 set() 作为值
```

### 为什么使用 `defaultdict`？

使用 `defaultdict` 可以避免在访问字典时出现 `KeyError`，并且可以简化代码，特别是在你需要初始化字典中的每个键为特定类型的值时。

### 注意事项

- `defaultdict` 的默认值是通过工厂函数生成的，每次访问不存在的键时都会调用这个函数，所以如果这个函数有副作用，它可能会影响你的程序。
- 使用 `defaultdict` 时，要确保提供的工厂函数在每次调用时都能生成相同的初始值。