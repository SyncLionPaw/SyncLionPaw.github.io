import time

def lpt_schedule(tasks, m):
    """
    LPT (Longest Processing Time) 贪心算法
    tasks: 任务时间列表
    m: 机器数量
    """
    # 按任务处理时间降序排序
    sorted_tasks = sorted(tasks, reverse=True)

    # 初始化每台机器的工作时间
    machines = [0] * m
    # 记录每台机器分配的任务
    assignments = [[] for _ in range(m)]

    # 依次将任务分配给当前负载最小的机器
    for task in sorted_tasks:
        # 找到当前负载最小的机器
        min_load_machine = machines.index(min(machines))
        # 分配任务
        machines[min_load_machine] += task
        assignments[min_load_machine].append(task)

    return assignments, max(machines)

if __name__ == "__main__":
    # 使用与回溯算法相同的测试用例
    tasks = [1, 2, 3, 4, 5, 6, 7, 8, 12, 23, 34, 45]
    m = 4  # 机器数量

    start = time.time()
    assignments, makespan = lpt_schedule(tasks, m)
    end = time.time()

    print(assignments, f"cost {round(end - start, 6)} seconds")