import heapq
from typing import List

"""
想象一下你是个城市基建规划者，地图上有 n 座城市，它们按以 1 到 n 的次序编号。

给你整数 n 和一个数组 conections，其中 connections[i] = [xi, yi, costi] 表示将城市 xi 和城市 yi 连接所要的costi（连接是双向的）。

返回连接所有城市的最低成本，每对城市之间至少有一条路径。如果无法连接所有 n 个城市，返回 -1

该 最小成本 应该是所用全部连接成本的总和。
"""


class Solution:
    def minimumCost(self, n: int, connections: List[List[int]]) -> int:
        u = Ufs(n)
        ans = 0
        edges = [(c, b, a) for (a, b, c) in connections]
        heapq.heapify(edges)
        while edges:
            cost, x, y = heapq.heappop(edges)
            x -= 1
            y -= 1
            if u.is_same(x, y):
                continue
            ans += cost
            u.union(x, y)
        if u.c != 1:
            return -1
        return ans


class Ufs:
    def __init__(self, size):
        self.p = list(range(size))
        self.r = [1] * size
        self.c = size

    def find(self, x) -> int:
        if self.p[x] != x:
            self.p[x] = self.find(self.p[x])
        return self.p[x]

    def union(self, x, y):
        rx, ry = self.find(x), self.find(y)
        if rx == ry:
            return
        if rx <= ry:
            self.p[rx] = ry
            self.r[ry] += self.r[rx]
        else:
            self.p[ry] = rx
            self.r[rx] += self.r[ry]
        self.c -= 1

    def is_same(self, x, y):
        return self.find(x) == self.find(y)
