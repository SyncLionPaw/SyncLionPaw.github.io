from typing import Dict


class Node:
    def __init__(self, val, prv=None, nxt=None, owner=None, key=0):
        self.prev = prv
        self.next = nxt
        self.val = val
        self.owner = owner
        self.key = key


class DoubleLinkList:
    def __init__(self):
        root = Node(0, owner=self)
        root.prev = root.next = root
        self.root = root
        self.size = 0

    def remove(self, node: Node):
        if node is None:
            return None
        assert node is not self.root and node.owner is self
        p, n = node.prev, node.next
        p.next = n
        n.prev = p
        node.prev = node.next = None
        self.size -= 1
        return node

    def insert_after(self, node: Node, at: Node):
        assert node is not None and at is not None
        nxt = at.next
        at.next = node
        node.prev = at
        node.next = nxt
        nxt.prev = node
        self.size += 1

    def insert_before(self, node: Node, at: Node):
        self.insert_after(node, at.prev)

    def insert_head(self, node):
        self.insert_after(node, self.root)

    def insert_tail(self, node: Node):
        self.insert_before(node, self.root)

    def remove_before_tail(self) -> Node:
        assert self.size > 0
        return self.remove(self.root.prev)


class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.link = DoubleLinkList()
        self.kn_map: Dict[int, Node] = {}

    def __len__(self) -> int:
        return self.link.size

    def Get(self, key) -> int:
        node = self.kn_map.get(key, None)
        if not node:
            return -1
        self.link.remove(node)
        self.link.insert_head(node)
        return node.val

    def Put(self, key, value):
        node = self.kn_map.get(key, None)
        if node is not None:
            node.val = value
            self.link.remove(node)
            self.link.insert_head(node)
            return
        newNode = Node(value, owner=self, key=key)

        if len(self) >= self.cap:
            drop = self.link.remove_before_tail()
            self.kn_map.pop(drop.key)

        self.kn_map[key] = newNode
        self.link.insert_head(newNode)
