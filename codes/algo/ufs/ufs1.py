# union found set impl
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


def test():
    ufs = Ufs(5)  # 创建一个包含5个元素的并查集

    ufs.union(0, 1)  # 合并集合（0, 1）
    ufs.union(2, 3)  # 合并集合（2, 3）

    # 查询元素0和元素1是否在同一集合中
    print(ufs.is_same(0, 1))  # 输出 True

    # 查询元素2和元素1是否在同一集合中
    print(ufs.is_same(1, 2))  # 输出 False

    ufs.union(1, 3)  # 合并集合（1, 3）

    # 查询元素0和元素3是否在同一集合中
    print(ufs.is_same(0, 3))

    # 查询元素0和元素4是否在同一集合中
    print(ufs.is_same(0, 4))  # 输出 False，表示0和4不在同一个集合中


test()
