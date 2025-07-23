from collections import defaultdict
from typing import Dict


class Node:
    def __init__(self, val, prv=None, nxt=None, owner=None, key=0):
        self.prev = prv
        self.next = nxt
        self.val = val
        self.owner = owner
        self.key = key


class FreqNode(Node):
    def __init__(self, val, prv=None, nxt=None, owner=None, key=0, freq=1):
        super().__init__(val=val, prv=prv, nxt=nxt, owner=owner, key=key)
        self.freq = freq


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
        newNode = Node(value, owner=self.link, key=key)

        if len(self) >= self.cap:
            drop = self.link.remove_before_tail()
            self.kn_map.pop(drop.key)

        self.kn_map[key] = newNode
        self.link.insert_head(newNode)


class LFUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.kn_map: Dict[int, Node] = {}
        self.links: Dict[int, DoubleLinkList] = defaultdict(DoubleLinkList)
        self.size = 0
        self.min_freq = 1

        self.size = 0  # maintain cache.size cause there were several links, for len, iterate costs O(n)

    def get(self, key: int) -> int:
        node: FreqNode = self.kn_map.get(key, None)
        if node is None:
            return -1
        link = self.links[node.freq]
        link.remove(node)

        # remove node from old link, may incr the min_freq
        if link.size == 0 and self.min_freq == node.freq:
            self.min_freq += 1

        node.freq += 1
        upper_link = self.links[node.freq]
        node.owner = upper_link
        upper_link.insert_head(node)

        return node.val

    def put(self, key: int, value: int) -> None:
        node: FreqNode = self.kn_map.get(key, None)
        if node is not None:
            link = self.links[node.freq]
            link.remove(node)
            if self.min_freq == node.freq and link.size == 0:
                self.min_freq += 1
            node.freq += 1
            node.val = value  # may change val
            new_link = self.links[node.freq]
            new_link.insert_head(node)
            node.owner = new_link
            return
        one_freq_link = self.links[1]
        new_node = FreqNode(key=key, val=value, owner=one_freq_link, freq=1)

        if self.size >= self.cap:
            old_link = self.links[self.min_freq]
            drop = old_link.remove_before_tail()
            self.kn_map.pop(drop.key)
            self.size -= 1
        one_freq_link.insert_head(new_node)
        self.kn_map[key] = new_node
        self.min_freq = 1
        self.size += 1
