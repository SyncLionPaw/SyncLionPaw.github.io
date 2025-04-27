# union found set impl
class Ufs:
    def __init__(self, size):
        self.parent = list(range(size))  # 初始化每个节点的父节点为自己

    def union(self, a, b) -> bool:
        """并 合并为一个集合"""
        root_a, root_b = self.find(a), self.find(b)
        if root_a == root_b:
            return True
        self.parent[root_a] = root_b
        return False

    def find(self, a):  # 路径压缩，不完全的
        """查a的根节点"""
        p = self.parent[a]
        while p != self.parent[p]:
            p = self.parent[p]
        self.parent[a] = p
        return p
