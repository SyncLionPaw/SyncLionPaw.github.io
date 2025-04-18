from typing import List

N = int(input())

pre = []

for _ in range(N):
    s = list(map(int, input().split()))
    pre.append(s)
# print(pre)

# a, b = map(int, input().split())
# print(a, b)


# 计算依赖的数量
def count(i: int, pre: List[List[int]]) -> int:
    prelist = pre[i]
    if prelist[0] == 0:
        return 1
    ans = 1
    for x in prelist[1:]:
        ans += count(x-1, pre)
    return ans

ans = count(0, pre)
print(ans)