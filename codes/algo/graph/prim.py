"""展示一种错误的kruskal算法的实现 错误的点在于什么时候不能加入所选的边 什么时候可以加入->引入并查集"""
from collections import defaultdict
from typing import Dict, List
import heapq

"""
WARNING: WRONG ANSWER !!!
"""

g: Dict[int, List[tuple[int, int]]] = defaultdict(list)

V, E = map(int, input().split())

priority_queue = []

for _ in range(E):
    src, dest, dist = map(int, input().split())
    g[src].append((dest, dist))
    g[dest].append((src, dist))
    priority_queue.append((dist, src, dest))

heapq.heapify(priority_queue)

print(g)
print(priority_queue)

visited = [False] * (V + 1)  # id start from 1
ans = []
path = []
cost = 0

while priority_queue:
    dist, src, dest = heapq.heappop(priority_queue)  # get minimam dist
    sin, din = visited[src], visited[dest]
    if sin and din:
        continue
    if not sin:
        ans.append(src)
        visited[src] = True
    if not din:
        ans.append(dest)
        visited[dest] = True
    path.append((src, dest))
    cost += dist

print(ans, cost, path)

"""
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
