# stress_test.py
import threading
from mockdb import DBClient
from client import send_random_requests


def client_task(i):
    print("task", i)
    client = DBClient()
    num_requests = 20
    send_random_requests(client, num_requests)


threads = []
for i in range(200):  # 测试100个并发客户端
    t = threading.Thread(target=client_task, args=(i,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()
