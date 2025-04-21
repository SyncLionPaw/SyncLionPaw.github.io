# https://www.acwing.com/problem/content/submission/1382/
k, n = map(int, input().split())

values = []
while len(values) < n:
    values += list(map(int, input().split()))


# 求装满容量target的最小物品数目
def snapsack(values, cap):
    n = len(values)
    f = [[1<<31] * (cap+1) for _ in range(n+1)]

    for i in range(n+1):
        f[i][0] = 0
    
    for i, v in enumerate(values):
        for j in range(cap+1):
            if v > j:
                f[i+1][j] = f[i][j]
            else:
                f[i+1][j] = min(f[i][j], f[i+1][j-v] + 1)  # 是完全背包
    # for line in f:
    #     print(line)
    return f


# 求装满容量target的最小物品数目,空间压缩到O(n)
def snapsack_space_optimized(values, cap):
    n = len(values)
    f = [1<<31] * (cap+1)

    for i in range(n+1):
        f[0] = 0
    
    for i, v in enumerate(values):
        for j in range(cap+1):
            if v <= j:
                f[j] = min(f[j], f[j-v] + 1)  # 是完全背包
    # for line in f:
    #     print(line)
    return f

f = snapsack_space_optimized(values, max(values) * k)
for j in range(len(f)):
    if j == len(f) - 1 or f[j+1] > k:
        break
print(j) 
