import networkx as nx
import matplotlib.pyplot as plt

# 定义图的顶点数和边的列表
vertices = 7  # 顶点数目
edges = [
    (1, 2, 1),
    (1, 3, 1),
    (1, 5, 2),
    (2, 6, 1),
    (2, 4, 2),
    (2, 3, 2),
    (3, 4, 1),
    (4, 5, 1),
    (5, 6, 2),
    (5, 7, 1),
    (6, 7, 1),
]

# 创建一个空的无向图
G = nx.Graph()

# 添加边到图中
for u, v, weight in edges:
    G.add_edge(u-1, v-1, weight=weight)

# 设置图形的布局
pos = nx.spring_layout(G)

# 绘制图形
plt.figure(figsize=(8, 6))

# 画图，节点标签为顶点编号，边标签为权重
nx.draw(
    G,
    pos,
    with_labels=True,
    node_size=500,
    node_color="lightblue",
    font_size=12,
    font_weight="bold",
)

edge_labels = nx.get_edge_attributes(G, "weight")

nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels)

plt.title("Graph Visualization")

plt.show()
