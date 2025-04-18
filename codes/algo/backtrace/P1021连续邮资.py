from math import inf


def max_covers(stamps: list, N: int):
    """stamps这个邮票组合,能够达到的连续邮资的最大值r"""
    stamp_type = len(stamps)
    max_value = stamps[-1] * N  # 理论能达到的最大面值
    min_use = [
        [inf] * (max_value + 1) for _ in range(stamp_type)
    ]  # min_use[i][j] 表示使用i前面的 邮票到达j面额，所需要的最小张数

    for j in range(max_value):
        if j * 1 > N:
            break
        min_use[0][j] = j

    for i in range(stamp_type):
        min_use[i][0] = 0

    for i in range(stamp_type):
        for j in range(1, max_value):
            cur_min = min_use[i][j]
            for t in range(N):
                if j < t * stamps[i]:
                    break
                cur_min = min(cur_min, t + min_use[i - 1][j - t * stamps[i]])
            min_use[i][j] = cur_min

    # for x in min_use:
    #     print(x)

    # r = [0] * stamp_type

    # for i in range(stamp_type):
    #     for j in range(max_value):
    #         if min_use[i][j + 1] > N and min_use[i][j] <= N:
    #             r[i] = j
    # return r
    max_reach = 0
    for j in range(1, max_value + 1):
        if min_use[stamp_type - 1][j] <= N:
            max_reach = j
        else:
            break
    return max_reach


# test
# ans = max_covers([1, 3], 3)
# print(ans)

N, K = map(int, input().split())

path = [1]
ans = 1
stamps = []


def backtrace(i):
    global ans, stamps
    r = max_covers(path, N)
    if i == K:
        if r >= ans:
            ans = r
            stamps = path.copy()
        return

    for v in range(path[-1] + 1, r + 2):
        path.append(v)
        backtrace(i+1)
        path.pop()
    

backtrace(1)
print(" ".join([str(x) for x in stamps]))
print(f"MAX={ans}")