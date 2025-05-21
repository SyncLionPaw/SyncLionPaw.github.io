# 差分数组

## 问题引入：

现在有一个数组 nums [1, 3, 5, 7, 0]
现在给你一个长度为2的数组a [2, 4] 表示把 nums 从下标 a[0] 到 a[1] 范围内的元素都加1。
求变化之后的数组的元素。

```py
nums = [1, 3, 5, 7, 0]
a = [2, 4]

for i in range(a[0], a[1]+1):
    nums[i] += 1
print(nums)
```
## 多次操作
若要进行多次操作呢？
例如 operators = [[1, 3], [2, 4], [1, 2], [2, 4]]

```py
operators = [[1, 3], [2, 4], [1, 2], [2, 4]]


nums = [1, 3, 5, 7, 0]

for start, end in operators:
    for i in range(start, end + 1):
        nums[i] += 1


print(nums)
```

若是频繁修改，每次修改的时间复杂度都是 O(n) 则不可接受。

## 解决办法
差分数组
先构建一个数组
我们先对 nums 数组构造一个 diff 差分数组，diff[i] 就是 nums[i] 和 nums[i-1] 之差：
```py
diff = [0] * len(nums)
# 构造差分数组
diff[0] = nums[0]
for i in range(1, len(nums)):
    diff[i] = nums[i] - nums[i - 1]
```
差分数组的每一项diff[i] = diff[i] - diff[i-1]

差分数组的每一项，都是 **我** 减去 **我前面的** 的差值。

通过这个 diff 差分数组是可以反推出原始数组 nums 的，代码逻辑如下：
```py
res = [0] * len(diff)
# 根据差分数组构造结果数组
res[0] = diff[0]
for i in range(1, len(diff)):
    res[i] = res[i - 1] + diff[i]
```
恢复的时候，记得要加上前面的。

如果想对nums的区间[i..j]进行+3操作，就只需要让diff[i]+3, 然后diff[j+1] -= 3

原理很简单，回想 diff 数组反推 nums 数组的过程，diff[i] += 3 意味着给 nums[i..] 所有的元素都加了 3，然后 diff[j+1] -= 3 又意味着对于 nums[j+1..] 所有元素再减 3，那综合起来，是不是就是对 nums[i..j] 中的所有元素都加 3 了？

只要花费 O(1) 的时间修改 diff 数组，就相当于给 nums 的整个区间做了修改。多次修改 diff，然后通过 diff 数组反推，即可得到 nums 修改后的结果。