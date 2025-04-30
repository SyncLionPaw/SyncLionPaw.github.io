from collections import defaultdict
import heapq
from typing import Dict, List


class Ufs:
    def __init__(self, size):
        self.parent = list(range(size))  # 初始化每个节点的父节点为自己
        self.cnt = [1] * size  # 记录每一个集合的元素数目

    def union(self, a, b) -> bool:
        """并 合并为一个集合，返回原先是不是一个集合的"""
        root_a, root_b = self.find(a), self.find(b)
        if root_a == root_b:
            return True
        # 不是一个集合的，进行合并
        if self.cnt[root_a] <= self.cnt[root_b]:
            self.parent[root_a] = root_b
            self.cnt[root_b] += self.cnt[root_a]
        else:
            self.parent[root_b] = root_a
            self.cnt[root_a] += self.cnt[root_b]
        return False

    def is_same(self, a, b):
        return self.find(a) == self.find(b)

    def find(self, a):  # 递归法，路径压缩
        """查a的根节点"""
        if self.parent[a] != a:
            self.parent[a] = self.find(self.parent[a])
        return self.parent[a]


g: Dict[int, List[tuple[int, int]]] = defaultdict(list)

V, E = map(int, input().split())

priority_queue = []

for _ in range(E):  # 把边转换成领接表保存
    src, dest, dist = map(int, input().split())
    src -= 1
    dest -= 1
    g[src].append((dest, dist))
    g[dest].append((src, dist))
    priority_queue.append((dist, src, dest))

heapq.heapify(priority_queue)

print(g)
print(priority_queue)

path = []
cost = 0

u = Ufs(V)

while priority_queue:
    dist, src, dest = heapq.heappop(priority_queue)  # get minimam dist
    if u.is_same(dest, src):
        continue
    path.append((src, dest, dist))
    cost += dist
    u.union(src, dest)

print("min cost", cost, "path", path)

""" 7顶点 11边
7 11
1 2 1
1 3 1
1 5 2
2 6 1
2 4 2
2 3 2
3 4 1
4 5 1
5 6 2
5 7 1
6 7 1
"""