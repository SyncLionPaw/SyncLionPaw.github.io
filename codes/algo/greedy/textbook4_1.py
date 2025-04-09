# 算法设计与分析第五版 算法实现题4-1

# 使用小根堆维护当前会场中，最后一个活动的结束时间
# 

import heapq
import pathlib
from typing import List

import matplotlib.pyplot as plt

def draw_intervals(intervals):
    # 为每个区间分配一个垂直位置（为了画图错开）
    y_positions = list(range(len(intervals)))

    # 创建图形
    fig, ax = plt.subplots(figsize=(10, 3))

    # 绘制每个区间为一条水平线
    for i, ((start, end), y) in enumerate(zip(intervals, y_positions)):
        ax.hlines(y, start, end, colors='skyblue', linewidth=6)
        ax.text(start, y + 0.2, f'{start}-{end}', fontsize=9)

    # 设置坐标轴
    ax.set_yticks(y_positions)
    ax.set_yticklabels([f'Activity {i+1}' for i in range(len(intervals))])
    ax.set_xlabel('Time')
    ax.set_title('activities')
    ax.grid(True, axis='x')
    plt.tight_layout()
    cwd = pathlib.Path(__file__)
    filepath = cwd.parent / "textbool4_1.png"

    plt.savefig(filepath)


def min_meeting_rooms(intervals: List[List[int]]):
    intervals.sort(key = lambda x:x[0])  # 按照开始时间升序排列
    heap = []
    for start, end in intervals:
        if heap and heap[0] <= start:
            heapq.heappop(heap)
        heapq.heappush(heap, end)
    return len(heap)

if __name__ == '__main__':
    intervals = [[1, 23], [12, 28], [25, 35], [27, 80], [36, 50]]
    draw_intervals(intervals)
    print(min_meeting_rooms(intervals))
