import copy
import time


tasks = [1, 2, 3, 4, 5, 6, 7, 8, 12, 23, 34, 45]
n = len(tasks)
m = 4  # machine

groups = [[] for _ in range(m)]
groups_sum = [0] * m

ans = 1e12
ans_case = None


def dfs(i: int):
    global ans, ans_case
    if i == n:
        if max(groups_sum) < ans:
            ans_case = copy.deepcopy(groups)
            ans = max(groups_sum)
        return
    for j in range(m):
        t = tasks[i]
        groups[j].append(t)
        groups_sum[j] += t
        dfs(i + 1)
        groups[j].pop()
        groups_sum[j] -= t
    return


start = time.time()
dfs(0)
end = time.time()

print(ans_case, ans, f"cost {round(end - start, 3)} seconds")
